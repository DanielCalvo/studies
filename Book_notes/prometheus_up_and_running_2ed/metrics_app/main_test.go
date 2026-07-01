package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

func TestHelloWorldRecordsSuccessMetrics(t *testing.T) {
	registry := prometheus.NewRegistry()
	application := newApp(registry, 0)
	application.randomFloat = func() float64 { return 0.5 }

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/", nil)

	application.helloWorld(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
	if got := testutil.ToFloat64(application.metrics.requests); got != 1 {
		t.Fatalf("expected 1 request, got %v", got)
	}
	if got := testutil.ToFloat64(application.metrics.exceptions); got != 0 {
		t.Fatalf("expected 0 exceptions, got %v", got)
	}
	if got := testutil.ToFloat64(application.metrics.sales); got != 0.5 {
		t.Fatalf("expected 0.5 sales euros, got %v", got)
	}
	if got := testutil.ToFloat64(application.metrics.inProgress); got != 0 {
		t.Fatalf("expected no requests in progress after handler, got %v", got)
	}
	if got := testutil.CollectAndCount(application.metrics.latency); got == 0 {
		t.Fatal("expected latency histogram samples")
	}
}

func TestHelloWorldRecordsForcedException(t *testing.T) {
	registry := prometheus.NewRegistry()
	application := newApp(registry, 0)
	application.randomFloat = func() float64 { return 0.5 }

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/?fail=true", nil)

	application.helloWorld(recorder, request)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf("expected status %d, got %d", http.StatusInternalServerError, recorder.Code)
	}
	if got := testutil.ToFloat64(application.metrics.requests); got != 1 {
		t.Fatalf("expected 1 request, got %v", got)
	}
	if got := testutil.ToFloat64(application.metrics.exceptions); got != 1 {
		t.Fatalf("expected 1 exception, got %v", got)
	}
	if got := testutil.ToFloat64(application.metrics.sales); got != 0 {
		t.Fatalf("expected no sales on failed request, got %v", got)
	}
}

func TestParseErrorRate(t *testing.T) {
	tests := map[string]struct {
		raw     string
		want    float64
		wantErr bool
	}{
		"default":  {raw: "", want: defaultErrorRate},
		"zero":     {raw: "0", want: 0},
		"one":      {raw: "1", want: 1},
		"fraction": {raw: "0.25", want: 0.25},
		"negative": {raw: "-0.1", wantErr: true},
		"too high": {raw: "1.1", wantErr: true},
		"not num":  {raw: "sometimes", wantErr: true},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := parseErrorRate(test.raw)
			if test.wantErr {
				if err == nil {
					t.Fatal("expected error")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != test.want {
				t.Fatalf("expected %v, got %v", test.want, got)
			}
		})
	}
}
