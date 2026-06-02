package main

import (
	"log"
	"net/http"
	"os"

	"alerts/internal"

	"battery-fleet-monitor/pkg/logger"
	"battery-fleet-monitor/pkg/natsx"
)

func main() {

	logg := logger.New("alerts")

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

	publisher := internal.NewPublisher(
		conn,
	)

	err = internal.StartSubscriber(
		conn,
		store,
		publisher,
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
		"alert service started",
		"port",
		8081,
	)

	err = http.ListenAndServe(
		":8081",
		mux,
	)

	if err != nil {
		log.Fatal(err)
	}
}
