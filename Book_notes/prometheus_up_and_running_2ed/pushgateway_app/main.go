package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

const (
	defaultGatewayURL   = "http://localhost:9092"
	defaultJobName      = "batch"
	defaultWorkDuration = time.Second
)

type config struct {
	gatewayURL   string
	jobName      string
	workDuration time.Duration
	failJob      bool
}

func main() {
	cfg, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := run(cfg); err != nil {
		log.Fatal(err)
	}
}

func run(cfg config) error {
	registry := prometheus.NewRegistry()

	duration := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "my_job_duration_seconds",
		Help: "Duration of my batch job in seconds.",
	})
	registry.MustRegister(duration)

	start := time.Now()
	jobErr := doWork(cfg)
	duration.Set(time.Since(start).Seconds())

	if jobErr == nil {
		lastSuccess := prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "my_job_last_success_seconds",
			Help: "Last time my batch job successfully finished.",
		})
		registry.MustRegister(lastSuccess)
		lastSuccess.SetToCurrentTime()
	}

	pushErr := push.New(cfg.gatewayURL, cfg.jobName).
		Gatherer(registry).
		Add()

	if jobErr != nil && pushErr != nil {
		return fmt.Errorf("batch failed: %w; also failed to push metrics: %v", jobErr, pushErr)
	}
	if jobErr != nil {
		return jobErr
	}
	if pushErr != nil {
		return fmt.Errorf("failed to push metrics: %w", pushErr)
	}

	log.Printf("pushed metrics to %s for job %q", cfg.gatewayURL, cfg.jobName)
	return nil
}

func doWork(cfg config) error {
	time.Sleep(cfg.workDuration)
	if cfg.failJob {
		return errors.New("simulated batch job failure")
	}
	return nil
}

func loadConfig() (config, error) {
	workDuration := defaultWorkDuration
	if raw := os.Getenv("WORK_DURATION_SECONDS"); raw != "" {
		seconds, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return config{}, fmt.Errorf("invalid WORK_DURATION_SECONDS: %w", err)
		}
		if seconds < 0 {
			return config{}, fmt.Errorf("WORK_DURATION_SECONDS must be non-negative, got %v", seconds)
		}
		workDuration = time.Duration(seconds * float64(time.Second))
	}

	return config{
		gatewayURL:   envOrDefault("PUSHGATEWAY_URL", defaultGatewayURL),
		jobName:      envOrDefault("JOB_NAME", defaultJobName),
		workDuration: workDuration,
		failJob:      os.Getenv("FAIL_JOB") == "true",
	}, nil
}

func envOrDefault(name, fallback string) string {
	value := os.Getenv(name)
	if value == "" {
		return fallback
	}
	return value
}
