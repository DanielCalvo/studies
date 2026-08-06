package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

const maxPayloadBytes = 1 << 20

type alertPayload struct {
	Alerts []alertNotification `json:"alerts"`
}

type alertNotification struct {
	Status      string            `json:"status"`
	Labels      map[string]string `json:"labels"`
	Annotations map[string]string `json:"annotations"`
	StartsAt    string            `json:"startsAt"`
	EndsAt      string            `json:"endsAt"`
	Fingerprint string            `json:"fingerprint"`
}

type displayedAlert struct {
	Fingerprint string
	Status      string
	Alertname   string
	Severity    string
	Environment string
	Summary     string
	StartsAt    string
	EndsAt      string
	UpdatedAt   time.Time
}

var alertStore = struct {
	sync.RWMutex
	alerts map[string]displayedAlert
}{alerts: make(map[string]displayedAlert)}

var pageTemplate = template.Must(template.New("alerts").Parse(`<!doctype html>
<html lang="en"><head><meta charset="utf-8"><meta http-equiv="refresh" content="5">
<title>Alert sink</title><style>
body{font:16px sans-serif;margin:2rem;background:#f7f7f7;color:#222} table{border-collapse:collapse;width:100%;background:white}
th,td{border:1px solid #ddd;padding:.6rem;text-align:left} th{background:#eee}
.firing{color:#b42318;font-weight:bold}.resolved{color:#067647;font-weight:bold}
</style></head><body><h1>Alert sink</h1>
<p>In-memory Alertmanager webhook state. Refreshes every five seconds; pod restarts clear this table.</p>
<table><thead><tr><th>Status</th><th>Alert</th><th>Summary</th><th>Severity</th><th>Environment</th><th>Started</th><th>Ended</th></tr></thead><tbody>
{{range .}}<tr><td class="{{.Status}}">{{.Status}}</td><td>{{.Alertname}}</td><td>{{.Summary}}</td><td>{{.Severity}}</td><td>{{.Environment}}</td><td>{{.StartsAt}}</td><td>{{.EndsAt}}</td></tr>{{else}}<tr><td colspan="7">No alerts received yet.</td></tr>{{end}}
</tbody></table></body></html>`))

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", showAlerts)
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = io.WriteString(w, "ok\n")
	})
	mux.HandleFunc("GET /api/alerts", listAlerts)
	mux.HandleFunc("POST /alerts", receiveAlerts)

	log.Println("alert sink listening on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func receiveAlerts(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(http.MaxBytesReader(w, r.Body, maxPayloadBytes))
	if err != nil {
		http.Error(w, "could not read payload", http.StatusBadRequest)
		return
	}

	var payload alertPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		http.Error(w, "payload must be Alertmanager JSON", http.StatusBadRequest)
		return
	}

	alertStore.Lock()
	for _, notification := range payload.Alerts {
		if notification.Fingerprint == "" {
			continue
		}
		status := notification.Status
		if status != "resolved" {
			status = "firing"
		}
		alertStore.alerts[notification.Fingerprint] = displayedAlert{
			Fingerprint: notification.Fingerprint,
			Status:      status,
			Alertname:   notification.Labels["alertname"],
			Severity:    notification.Labels["severity"],
			Environment: notification.Labels["environment"],
			Summary:     notification.Annotations["summary"],
			StartsAt:    notification.StartsAt,
			EndsAt:      notification.EndsAt,
			UpdatedAt:   time.Now(),
		}
	}
	alertStore.Unlock()

	log.Printf("received alert payload: %s", body)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintln(w, `{"status":"ok"}`)
}

func listAlerts(w http.ResponseWriter, _ *http.Request) {
	alerts := snapshotAlerts()
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(alerts)
}

func showAlerts(w http.ResponseWriter, _ *http.Request) {
	if err := pageTemplate.Execute(w, snapshotAlerts()); err != nil {
		http.Error(w, "could not render alerts", http.StatusInternalServerError)
	}
}

func snapshotAlerts() []displayedAlert {
	alertStore.RLock()
	alerts := make([]displayedAlert, 0, len(alertStore.alerts))
	for _, alert := range alertStore.alerts {
		alerts = append(alerts, alert)
	}
	alertStore.RUnlock()
	sort.Slice(alerts, func(i, j int) bool {
		if alerts[i].Status != alerts[j].Status {
			return alerts[i].Status == "firing"
		}
		return alerts[i].UpdatedAt.After(alerts[j].UpdatedAt)
	})
	return alerts
}
