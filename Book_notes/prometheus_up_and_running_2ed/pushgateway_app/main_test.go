package main

import (
	"testing"
	"time"
)

func TestLoadConfigDefaults(t *testing.T) {
	t.Setenv("PUSHGATEWAY_URL", "")
	t.Setenv("JOB_NAME", "")
	t.Setenv("WORK_DURATION_SECONDS", "")
	t.Setenv("FAIL_JOB", "")

	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.gatewayURL != defaultGatewayURL {
		t.Fatalf("expected gateway %q, got %q", defaultGatewayURL, cfg.gatewayURL)
	}
	if cfg.jobName != defaultJobName {
		t.Fatalf("expected job %q, got %q", defaultJobName, cfg.jobName)
	}
	if cfg.workDuration != defaultWorkDuration {
		t.Fatalf("expected duration %s, got %s", defaultWorkDuration, cfg.workDuration)
	}
	if cfg.failJob {
		t.Fatal("expected failJob to default to false")
	}
}

func TestLoadConfigFromEnvironment(t *testing.T) {
	t.Setenv("PUSHGATEWAY_URL", "http://pushgateway:9091")
	t.Setenv("JOB_NAME", "nightly")
	t.Setenv("WORK_DURATION_SECONDS", "0.25")
	t.Setenv("FAIL_JOB", "true")

	cfg, err := loadConfig()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if cfg.gatewayURL != "http://pushgateway:9091" {
		t.Fatalf("unexpected gateway %q", cfg.gatewayURL)
	}
	if cfg.jobName != "nightly" {
		t.Fatalf("unexpected job %q", cfg.jobName)
	}
	if cfg.workDuration != 250*time.Millisecond {
		t.Fatalf("expected 250ms, got %s", cfg.workDuration)
	}
	if !cfg.failJob {
		t.Fatal("expected failJob to be true")
	}
}

func TestLoadConfigRejectsInvalidDuration(t *testing.T) {
	t.Setenv("WORK_DURATION_SECONDS", "-1")

	if _, err := loadConfig(); err == nil {
		t.Fatal("expected error")
	}
}
