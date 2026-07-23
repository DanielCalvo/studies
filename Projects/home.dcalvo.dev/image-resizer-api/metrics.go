package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metrics struct {
	registry       *prometheus.Registry
	httpRequests   *prometheus.CounterVec
	httpDuration   *prometheus.HistogramVec
	httpInFlight   *prometheus.GaugeVec
	stageDuration  *prometheus.HistogramVec
	resizeInFlight prometheus.Gauge
	inputBytes     prometheus.Histogram
	outputBytes    prometheus.Histogram
	inputPixels    prometheus.Histogram
	outputPixels   prometheus.Histogram
	rejections     *prometheus.CounterVec
}

func newMetrics() *metrics {
	requestDurationBuckets := []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10, 20, 30}
	byteBuckets := prometheus.ExponentialBuckets(1<<10, 4, 8)
	pixelBuckets := prometheus.ExponentialBuckets(100_000, 4, 6)

	m := &metrics{
		registry: prometheus.NewRegistry(),
		httpRequests: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "image_resizer",
			Subsystem: "http",
			Name:      "requests_total",
			Help:      "Total number of resize HTTP requests.",
		}, []string{"method", "route", "status", "outcome"}),
		httpDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Subsystem: "http",
			Name:      "request_duration_seconds",
			Help:      "Duration of resize HTTP requests in seconds.",
			Buckets:   requestDurationBuckets,
		}, []string{"method", "route", "outcome"}),
		httpInFlight: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "image_resizer",
			Subsystem: "http",
			Name:      "requests_in_flight",
			Help:      "Current number of resize HTTP requests being handled.",
		}, []string{"method", "route"}),
		stageDuration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Name:      "processing_stage_duration_seconds",
			Help:      "Duration of image processing stages in seconds.",
			Buckets:   requestDurationBuckets,
		}, []string{"stage"}),
		resizeInFlight: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "image_resizer",
			Name:      "resize_operations_in_flight",
			Help:      "Current number of images being decoded, resized, or encoded.",
		}),
		inputBytes: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Name:      "input_size_bytes",
			Help:      "Encoded size of accepted input JPEG images in bytes.",
			Buckets:   byteBuckets,
		}),
		outputBytes: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Name:      "output_size_bytes",
			Help:      "Encoded size of successfully resized JPEG images in bytes.",
			Buckets:   byteBuckets,
		}),
		inputPixels: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Name:      "input_pixels",
			Help:      "Pixel count of accepted input JPEG images.",
			Buckets:   pixelBuckets,
		}),
		outputPixels: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: "image_resizer",
			Name:      "output_pixels",
			Help:      "Pixel count of successfully resized JPEG images.",
			Buckets:   pixelBuckets,
		}),
		rejections: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "image_resizer",
			Name:      "rejections_total",
			Help:      "Total number of rejected resize requests by stable reason.",
		}, []string{"reason"}),
	}

	m.registry.MustRegister(
		m.httpRequests,
		m.httpDuration,
		m.httpInFlight,
		m.stageDuration,
		m.resizeInFlight,
		m.inputBytes,
		m.outputBytes,
		m.inputPixels,
		m.outputPixels,
		m.rejections,
		prometheus.NewGoCollector(),
		prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}),
	)
	return m
}

func (m *metrics) handler() http.Handler {
	return promhttp.HandlerFor(m.registry, promhttp.HandlerOpts{EnableOpenMetrics: true})
}

func (m *metrics) startRequest(method, route string) {
	m.httpInFlight.WithLabelValues(metricMethod(method), route).Inc()
}

func (m *metrics) finishRequest(method, route string, status int, duration time.Duration) {
	method = metricMethod(method)
	outcome := requestOutcome(status)
	m.httpInFlight.WithLabelValues(method, route).Dec()
	m.httpRequests.WithLabelValues(method, route, strconv.Itoa(status), outcome).Inc()
	m.httpDuration.WithLabelValues(method, route, outcome).Observe(duration.Seconds())
}

func metricMethod(method string) string {
	if method == http.MethodPost {
		return method
	}
	return "other"
}

func (m *metrics) observeStage(stage string, duration time.Duration) {
	m.stageDuration.WithLabelValues(stage).Observe(duration.Seconds())
}

func requestOutcome(status int) string {
	switch {
	case status >= 500:
		return "failed"
	case status >= 400:
		return "rejected"
	default:
		return "succeeded"
	}
}
