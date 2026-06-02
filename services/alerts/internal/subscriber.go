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
	pub *Publisher,
	logg *slog.Logger,
) error {

	return natsx.Subscribe[models.BatteryTelemetry](
		conn,
		events.BatteryTelemetrySubject,
		func(event models.BatteryTelemetry) {

			alerts := Evaluate(event)

			for _, alert := range alerts {

				store.Add(alert)

				err := pub.Publish(alert)

				if err != nil {

					logg.Error(
						"failed publishing alert",
						"error",
						err,
					)

					continue
				}

				logg.Warn(
					"battery alert",
					"battery_id",
					alert.BatteryID,
					"severity",
					alert.Severity,
					"message",
					alert.Message,
				)
			}
		},
	)
}
