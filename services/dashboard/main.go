package main

import (
	"log"
	"net/http"
	"os"

	"dashboard/internal"

	"battery-fleet-monitor/pkg/logger"
)

func main() {

	logg := logger.New("dashboard")

	telemetryURL := os.Getenv(
		"TELEMETRY_URL",
	)

	if telemetryURL == "" {
		telemetryURL = "http://localhost:8080"
	}

	alertURL := os.Getenv(
		"ALERT_URL",
	)

	if alertURL == "" {
		alertURL = "http://localhost:8081"
	}

	client := internal.NewClient(
		telemetryURL,
		alertURL,
	)

	mux := http.NewServeMux()

	internal.RegisterHandlers(
		mux,
		client,
	)

	logg.Info(
		"dashboard service started",
		"port",
		8082,
	)

	err := http.ListenAndServe(
		":8082",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
