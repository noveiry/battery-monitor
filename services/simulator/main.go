package main

import (
	"log"
	"math/rand"
	"os"
	"time"

	"battery-fleet-monitor/pkg/logger"
	"battery-fleet-monitor/pkg/natsx"

	"simulator/internal"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	logg := logger.New("simulator")

	natsURL := os.Getenv("NATS_URL")

	if natsURL == "" {
		natsURL = "nats://localhost:4222"
	}

	conn, err := natsx.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	batteryID := os.Getenv("BATTERY_ID")

	if batteryID == "" {
		batteryID = internal.NewRandomID()
	}

	generator := internal.NewGenerator(
		batteryID,
	)

	publisher := internal.NewPublisher(
		conn,
	)

	logg.Info(
		"battery simulator started",
		"battery_id",
		batteryID,
	)

	ticker := time.NewTicker(
		2 * time.Second,
	)

	defer ticker.Stop()

	for range ticker.C {

		event := generator.Next()

		err := publisher.Publish(
			event,
		)

		if err != nil {

			logg.Error(
				"publish failed",
				"error",
				err,
			)

			continue
		}

		logg.Info(
			"telemetry published",
			"id",
			event.ID,
			"charge",
			event.Charge,
			"temp",
			event.Temperature,
			"health",
			event.Health,
			"status",
			event.Status,
		)
	}
}