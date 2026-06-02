package internal

import (
	"log/slog"

	"battery-fleet-monitor/pkg/events"
	"battery-fleet-monitor/pkg/models"
	"battery-fleet-monitor/pkg/natsx"

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
				"temp",
				event.Temperature,
			)
		},
	)
}
