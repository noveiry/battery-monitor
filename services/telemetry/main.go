package main

import (
	"battery-monitor/pkg/logger"
	"battery-monitor/pkg/natsx"
	"battery-monitor/services/telemetry/internal"
	"log"
	"net/http"
	"os"
)

func main() {

	logg := logger.New("telemetry")

	natsURL := os.Getenv("NATS_URL")

	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	conn, err := natsx.Connect(
		natsURL,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	store := internal.NewStore()

	err = internal.StartSubscriber(
		conn,
		store,
		logg,
	)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	internal.RegisterHandlers(
		mux,
		store,
	)

	logg.Info(
		"telemetry service started",
		"port",
		8080,
	)

	err = http.ListenAndServe(
		":8080",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
