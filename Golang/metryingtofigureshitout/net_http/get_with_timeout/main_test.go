package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetWithTimeout(t *testing.T) {
	t.Run("Do HEAD should not time out", func(t *testing.T) {
		srv := MakeNewTestServerWithDelay(20 * time.Millisecond)
		defer srv.Close()
		_, err := HeadUrlWithTimeout(srv.URL, 50*time.Millisecond)
		if err != nil {
			t.Errorf("HTTP request timeout when it should not have timed out")
		}
	})

	t.Run("Do HEAD request that should time out", func(t *testing.T) {
		srv := MakeNewTestServerWithDelay(20 * time.Millisecond)
		defer srv.Close()
		_, err := HeadUrlWithTimeout(srv.URL, 10*time.Millisecond)
		if err == nil {
			t.Errorf("Expected a timeout error, received no error")
		}
	})
}

func MakeNewTestServerWithDelay(delay time.Duration) *httptest.Server {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
	return srv
}
