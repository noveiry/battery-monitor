package internal

import (
	"time"

	"battery-monitor/pkg/models"
)

func Evaluate(
	b models.BatteryTelemetry,
) []models.BatteryAlert {

	var alerts []models.BatteryAlert

	if b.Charge < 10 {

		alerts = append(alerts,
			models.BatteryAlert{
				BatteryID: b.ID,
				Severity:  "critical",
				Message:   "battery charge below 10%",
				Time:      time.Now().UTC(),
			},
		)
	}

	if b.Temperature > 50 {

		alerts = append(alerts,
			models.BatteryAlert{
				BatteryID: b.ID,
				Severity:  "critical",
				Message:   "battery overheating",
				Time:      time.Now().UTC(),
			},
		)
	}

	if b.Health < 70 {

		alerts = append(alerts,
			models.BatteryAlert{
				BatteryID: b.ID,
				Severity:  "warning",
				Message:   "battery health degraded",
				Time:      time.Now().UTC(),
			},
		)
	}

	return alerts
}
