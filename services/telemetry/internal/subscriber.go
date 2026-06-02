package internal

import (
	"log/slog"

	"battery-monitor/pkg/events"
	"battery-monitor/pkg/models"
	"battery-monitor/pkg/natsx"

	"github.com/nats-io/nats.go"
)

func StartSubscriber(
	conn *nats.Conn,
	store *Store,
	logg *slog.Logger,
) error {

	return natsx.Subscribe[models.BatteryTelemetry](
		conn,
		events.BatteryTelemetrySubject,
		func(event models.BatteryTelemetry) {

			store.Upsert(event)

			logg.Info(
				"telemetry received",
				"id",
				event.ID,
				"charge",
				event.Charge,
				"temperature",
				event.Temperature,
			)
		},
	)
}
