package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
	"syscall"
	"time"
	"unicode/utf8"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	otelprometheus "go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

const (
	listenAddress          = ":8080"
	namesFile              = "names.txt"
	instrumentationName    = "ai-incremental-metrics"
	instrumentationVersion = "1.0.0"
)

// The Go HTTP server handles requests concurrently. This mutex ensures that
// two requests do not read or modify the names file at the same time.
var namesFileMu sync.Mutex

// nameOperations counts successful CRUD operations. The instrument is created
// once during startup and can then be used safely by concurrent HTTP handlers.
var nameOperations metric.Int64Counter

// readNameOperations is an independently requested handle for the same logical
// instrument. Its registration is identical to nameOperations.
var readNameOperations metric.Int64Counter

// createDuration records the distribution of time spent in create requests.
var createDuration metric.Float64Histogram

// activeRequests tracks how many HTTP requests are currently being processed.
var activeRequests metric.Int64UpDownCounter

// storedNames reports the absolute number of names found during collection.
var storedNames metric.Int64ObservableUpDownCounter

// longestNameLength reports a non-additive current value during collection.
var longestNameLength metric.Int64ObservableGauge

// majorPageFaults observes an operating-system cumulative value during collection.
var majorPageFaults metric.Int64ObservableCounter

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "server failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// The Prometheus exporter is also an OpenTelemetry metric reader. It
	// collects metrics when Prometheus (or curl) scrapes the HTTP endpoint,
	// rather than collecting on an internal timer.
	prometheusReader, err := otelprometheus.New()
	if err != nil {
		return fmt.Errorf("create Prometheus metrics reader: %w", err)
	}

	// A View changes how matching instruments are aggregated by the SDK. This
	// one selects useful duration buckets for our fast local file operation.
	// The boundaries use seconds because that is the histogram's unit.
	createDurationView := sdkmetric.NewView(
		sdkmetric.Instrument{Name: "names.create.duration"},
		sdkmetric.Stream{
			Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
				Boundaries: []float64{
					0.001, // 1 ms
					0.005, // 5 ms
					0.010, // 10 ms
					0.025, // 25 ms
					0.050, // 50 ms
					0.100, // 100 ms
					0.250, // 250 ms
					0.500, // 500 ms
					1.000, // 1 s
				},
			},
		},
	)

	// Both of these Views match the same counter. The first preserves its
	// existing per-operation stream. The second removes every metric attribute
	// and gives the resulting rolled-up stream a distinct name.
	operationCriteria := sdkmetric.Instrument{Name: "names.operations"}
	operationDetailView := sdkmetric.NewView(
		operationCriteria,
		sdkmetric.Stream{},
	)
	allOperationsView := sdkmetric.NewView(
		operationCriteria,
		sdkmetric.Stream{
			Name:            "names.operations.all",
			Description:     "Total successful name operations across all operation types",
			AttributeFilter: attribute.NewAllowKeysFilter(), //   This setting removes every metric attribute from the rolled-up stream
		},
	)

	// A Resource identifies the entity producing all telemetry from this
	// provider. Merge our service identity with the SDK's default resource so
	// useful metadata such as the telemetry SDK name and version is preserved.
	serviceResource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("names-service"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return fmt.Errorf("create service resource: %w", err)
	}

	// The MeterProvider owns the metrics pipeline. Connecting the reader here
	// gives measurements recorded through this provider a path to the exporter.
	// The View configures aggregation, while the Resource identifies the entity
	// producing every metric from this provider.
	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(prometheusReader),
		sdkmetric.WithView(
			createDurationView,
			operationDetailView,
			allOperationsView,
		),
		sdkmetric.WithResource(serviceResource),
	)
	otel.SetMeterProvider(provider)
	defer func() {
		if err := provider.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "shut down metrics provider: %v\n", err)
		}
	}()

	// A Meter creates metric instruments such as counters and histograms. Its
	// name, version, and schema URL identify the instrumentation scope: the code
	// that defines those instruments. This is separate from the service
	// Resource, which identifies the entity running that code.
	meter := otel.Meter(
		instrumentationName,
		metric.WithInstrumentationVersion(instrumentationVersion),
		metric.WithSchemaURL(semconv.SchemaURL),
	)

	// A counter records a total that can only increase. Attributes supplied
	// when recording will divide this total into separate CRUD operation series.
	nameOperations, err = meter.Int64Counter(
		"names.operations",
		metric.WithDescription("Number of successful name operations"),
	)
	if err != nil {
		return fmt.Errorf("create name operations counter: %w", err)
	}

	// A separate application component can independently request an identical
	// instrument. Because the Meter, name, type, unit, and description match,
	// the SDK treats both handles as the same logical instrument.
	//
	// This second handle offers no practical benefit in this small program
	// because every handler can already share nameOperations. It exists only to
	// illustrate how identical registrations from independent components are
	// combined; without that learning goal, this service should use one handle.
	readNameOperations, err = meter.Int64Counter(
		"names.operations",
		metric.WithDescription("Number of successful name operations"),
	)
	if err != nil {
		return fmt.Errorf("create read name operations counter: %w", err)
	}

	// A histogram records a distribution of measurements rather than a
	// continuously increasing total. We use seconds as the duration unit.
	createDuration, err = meter.Float64Histogram(
		"names.create.duration",
		metric.WithUnit("s"),
		metric.WithDescription("Duration of create-name operations"),
	)
	if err != nil {
		return fmt.Errorf("create name duration histogram: %w", err)
	}

	// An up/down counter records a value that can increase and decrease. Each
	// request will add one when it begins and subtract one when it finishes.
	activeRequests, err = meter.Int64UpDownCounter(
		"http.server.active_requests",
		metric.WithDescription("Number of HTTP requests currently being processed"),
	)
	if err != nil {
		return fmt.Errorf("create active requests counter: %w", err)
	}

	// An observable instrument runs its callback when the reader collects
	// metrics. The callback reports the current absolute value rather than
	// recording each create and delete as a change.
	storedNames, err = meter.Int64ObservableUpDownCounter(
		"names.stored",
		metric.WithUnit("{name}"),
		metric.WithDescription("Current number of names stored in the file"),
		metric.WithInt64Callback(func(_ context.Context, observer metric.Int64Observer) error {
			namesFileMu.Lock()
			defer namesFileMu.Unlock()

			names, err := loadNames()
			if err != nil {
				return fmt.Errorf("load names for metric callback: %w", err)
			}

			observer.Observe(int64(len(names)))
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("create stored names counter: %w", err)
	}

	// A gauge reports a current value that should not be summed across sources.
	// The longest lengths from two service instances should be compared with
	// max, not added together.
	longestNameLength, err = meter.Int64ObservableGauge(
		"names.longest_length",
		metric.WithUnit("{character}"),
		metric.WithDescription("Length of the longest currently stored name"),
		metric.WithInt64Callback(func(_ context.Context, observer metric.Int64Observer) error {
			namesFileMu.Lock()
			defer namesFileMu.Unlock()

			names, err := loadNames()
			if err != nil {
				return fmt.Errorf("load names for longest-name callback: %w", err)
			}

			longestLength := 0
			for _, name := range names {
				longestLength = max(longestLength, utf8.RuneCountInString(name))
			}

			observer.Observe(int64(longestLength))
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("create longest name gauge: %w", err)
	}

	// An observable counter reports an absolute cumulative value that must not
	// decrease. Linux already maintains the process's major-page-fault total,
	// so the SDK reads it during a scrape instead of our code calling Add.
	majorPageFaults, err = meter.Int64ObservableCounter(
		"process.major_page_faults",
		metric.WithUnit("{fault}"),
		metric.WithDescription("Major page faults requiring I/O since the process started"),
		metric.WithInt64Callback(func(_ context.Context, observer metric.Int64Observer) error {
			var usage syscall.Rusage
			if err := syscall.Getrusage(syscall.RUSAGE_SELF, &usage); err != nil {
				return fmt.Errorf("read process resource usage: %w", err)
			}

			observer.Observe(usage.Majflt)
			return nil
		}),
	)
	if err != nil {
		return fmt.Errorf("create major page faults counter: %w", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /names", createName)
	mux.HandleFunc("GET /names", listNames)
	mux.HandleFunc("GET /names/{name}", readName)
	mux.HandleFunc("PUT /names/{name}", updateName)
	mux.HandleFunc("DELETE /names/{name}", deleteName)

	// promhttp serves the collectors registered with Prometheus's default
	// registry. A request to this endpoint makes the OpenTelemetry Prometheus
	// reader collect synchronous aggregates and invoke observable callbacks.
	mux.Handle("GET /metrics", promhttp.Handler())

	fmt.Printf("Names service listening on http://localhost%s\n", listenAddress)
	return http.ListenAndServe(listenAddress, trackActiveRequests(mux))
}

// trackActiveRequests wraps every HTTP request, regardless of which route it
// uses or whether it succeeds. This is application middleware, not automatic
// HTTP instrumentation supplied by OpenTelemetry.
func trackActiveRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		activeRequests.Add(r.Context(), 1)
		defer func() {
			// Deferring the subtraction ensures that early returns and error
			// responses do not leave the active-request count permanently high.
			activeRequests.Add(r.Context(), -1)
		}()

		next.ServeHTTP(w, r)
	})
}

// createName reads a name directly from the request body and appends it as one
// line in the file.
func createName(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	defer func() {
		// Record runs when the handler finishes, including when it returns early
		// because of an error. The SDK places this observation into histogram
		// buckets and keeps a count and sum for the complete distribution.
		createDuration.Record(r.Context(), time.Since(start).Seconds())
	}()

	name, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	if err := appendName(strings.TrimSpace(string(name))); err != nil {
		http.Error(w, "could not save name", http.StatusInternalServerError)
		return
	}

	// The operation attribute creates the "create" dimension of this counter.
	// Its value comes from a fixed set rather than from user-provided names.
	nameOperations.Add(
		r.Context(),
		1,
		metric.WithAttributes(attribute.String("operation", "create")),
	)

	w.WriteHeader(http.StatusCreated)
}

// listNames returns the complete file, where every line contains one name.
func listNames(w http.ResponseWriter, r *http.Request) {
	namesFileMu.Lock()
	defer namesFileMu.Unlock()

	contents, err := os.ReadFile(namesFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			nameOperations.Add(
				r.Context(),
				1,
				metric.WithAttributes(attribute.String("operation", "list")),
			)
			return
		}

		http.Error(w, "could not read names", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	nameOperations.Add(
		r.Context(),
		1,
		metric.WithAttributes(attribute.String("operation", "list")),
	)
	_, _ = w.Write(contents)
}

// readName checks whether the name in the URL is present in the file.
func readName(w http.ResponseWriter, r *http.Request) {
	namesFileMu.Lock()
	defer namesFileMu.Unlock()

	names, err := loadNames()
	if err != nil {
		http.Error(w, "could not read names", http.StatusInternalServerError)
		return
	}

	requestedName := r.PathValue("name")
	for _, name := range names {
		if name == requestedName {
			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			// This second handle contributes to the same aggregation used by
			// nameOperations; it does not create another Prometheus metric.
			readNameOperations.Add(
				r.Context(),
				1,
				metric.WithAttributes(attribute.String("operation", "read")),
			)
			_, _ = fmt.Fprintln(w, name)
			return
		}
	}

	http.Error(w, "name not found", http.StatusNotFound)
}

// updateName replaces every matching name with the name supplied in the
// request body.
func updateName(w http.ResponseWriter, r *http.Request) {
	replacement, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "could not read request body", http.StatusBadRequest)
		return
	}

	namesFileMu.Lock()
	defer namesFileMu.Unlock()

	names, err := loadNames()
	if err != nil {
		http.Error(w, "could not read names", http.StatusInternalServerError)
		return
	}

	found := false
	for index, name := range names {
		if name == r.PathValue("name") {
			names[index] = strings.TrimSpace(string(replacement))
			found = true
		}
	}

	if !found {
		http.Error(w, "name not found", http.StatusNotFound)
		return
	}

	if err := saveNames(names); err != nil {
		http.Error(w, "could not update name", http.StatusInternalServerError)
		return
	}

	nameOperations.Add(
		r.Context(),
		1,
		metric.WithAttributes(attribute.String("operation", "update")),
	)
	w.WriteHeader(http.StatusNoContent)
}

// deleteName removes every occurrence of the requested name from the file.
func deleteName(w http.ResponseWriter, r *http.Request) {
	namesFileMu.Lock()
	defer namesFileMu.Unlock()

	names, err := loadNames()
	if err != nil {
		http.Error(w, "could not read names", http.StatusInternalServerError)
		return
	}

	requestedName := r.PathValue("name")
	remainingNames := make([]string, 0, len(names))
	found := false
	for _, name := range names {
		if name == requestedName {
			found = true
			continue
		}
		remainingNames = append(remainingNames, name)
	}

	if !found {
		http.Error(w, "name not found", http.StatusNotFound)
		return
	}

	if err := saveNames(remainingNames); err != nil {
		http.Error(w, "could not delete name", http.StatusInternalServerError)
		return
	}

	nameOperations.Add(
		r.Context(),
		1,
		metric.WithAttributes(attribute.String("operation", "delete")),
	)
	w.WriteHeader(http.StatusNoContent)
}

func appendName(name string) error {
	namesFileMu.Lock()
	defer namesFileMu.Unlock()

	file, err := os.OpenFile(namesFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o600)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, name)
	return err
}

// loadNames expects its caller to hold namesFileMu.
func loadNames() ([]string, error) {
	contents, err := os.ReadFile(namesFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	}

	text := strings.TrimSuffix(string(contents), "\n")
	if text == "" {
		return nil, nil
	}

	return strings.Split(text, "\n"), nil
}

// saveNames expects its caller to hold namesFileMu.
func saveNames(names []string) error {
	contents := strings.Join(names, "\n")
	if len(names) > 0 {
		contents += "\n"
	}

	return os.WriteFile(namesFile, []byte(contents), 0o600)
}
