package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultMetricsAddr = ":8000"
	defaultAppAddr     = ":8001"
	defaultErrorRate   = 0.2
)

type appMetrics struct {
	requests   prometheus.Counter
	exceptions prometheus.Counter
	sales      prometheus.Counter
	inProgress prometheus.Gauge
	lastServed prometheus.Gauge
	latency    prometheus.Histogram
}

type app struct {
	metrics     appMetrics
	errorRate   float64
	randomFloat func() float64
}

func newApp(registerer prometheus.Registerer, errorRate float64) *app {
	factory := promauto.With(registerer)

	factory.NewGaugeFunc(prometheus.GaugeOpts{
		Name: "time_seconds",
		Help: "The current Unix time.",
	}, func() float64 {
		return float64(time.Now().UnixNano()) / float64(time.Second)
	})

	return &app{
		metrics: appMetrics{
			requests: factory.NewCounter(prometheus.CounterOpts{
				Name: "hello_worlds_total",
				Help: "Hello Worlds requested.",
			}),
			exceptions: factory.NewCounter(prometheus.CounterOpts{
				Name: "hello_world_exceptions_total",
				Help: "Exceptions serving Hello World.",
			}),
			sales: factory.NewCounter(prometheus.CounterOpts{
				Name: "hello_world_sales_euro_total",
				Help: "Euros made serving Hello World.",
			}),
			inProgress: factory.NewGauge(prometheus.GaugeOpts{
				Name: "hello_worlds_inprogress",
				Help: "Number of Hello Worlds in progress.",
			}),
			lastServed: factory.NewGauge(prometheus.GaugeOpts{
				Name: "hello_world_last_time_seconds",
				Help: "The last time a Hello World was served.",
			}),
			latency: factory.NewHistogram(prometheus.HistogramOpts{
				Name:    "hello_world_latency_seconds",
				Help:    "Time for a request Hello World.",
				Buckets: prometheus.DefBuckets,
			}),
		},
		errorRate:   errorRate,
		randomFloat: rand.Float64,
	}
}

func (a *app) helloWorld(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	a.metrics.inProgress.Inc()
	defer a.metrics.inProgress.Dec()
	defer func() {
		a.metrics.latency.Observe(time.Since(start).Seconds())
		a.metrics.lastServed.SetToCurrentTime()
	}()

	a.metrics.requests.Inc()

	if a.shouldFail(r) {
		a.metrics.exceptions.Inc()
		http.Error(w, "simulated Hello World failure", http.StatusInternalServerError)
		return
	}

	euros := a.randomFloat()
	a.metrics.sales.Add(euros)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintf(w, "Hello World for %.4f euros.\n", euros)
}

func (a *app) shouldFail(r *http.Request) bool {
	if r.URL.Query().Get("fail") == "true" {
		return true
	}
	return a.randomFloat() < a.errorRate
}

func main() {
	errorRate, err := parseErrorRate(os.Getenv("ERROR_RATE"))
	if err != nil {
		log.Fatalf("invalid ERROR_RATE: %v", err)
	}

	application := newApp(prometheus.DefaultRegisterer, errorRate)

	metricsMux := http.NewServeMux()
	metricsMux.Handle("/metrics", promhttp.Handler())
	metricsMux.Handle("/", promhttp.Handler())

	go func() {
		log.Printf("serving Prometheus metrics on %s", defaultMetricsAddr)
		if err := http.ListenAndServe(defaultMetricsAddr, metricsMux); err != nil {
			log.Fatalf("metrics server failed: %v", err)
		}
	}()

	appMux := http.NewServeMux()
	appMux.HandleFunc("/", application.helloWorld)

	log.Printf("serving Hello World app on %s", defaultAppAddr)
	if err := http.ListenAndServe(defaultAppAddr, appMux); err != nil {
		log.Fatalf("app server failed: %v", err)
	}
}

func parseErrorRate(raw string) (float64, error) {
	if raw == "" {
		return defaultErrorRate, nil
	}

	rate, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		return 0, err
	}
	if rate < 0 || rate > 1 {
		return 0, fmt.Errorf("must be between 0 and 1, got %v", rate)
	}
	return rate, nil
}
