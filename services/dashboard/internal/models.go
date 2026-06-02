package internal

import "battery-monitor/pkg/models"

type FleetResponse struct {
	Stats     any                       `json:"stats"`
	Batteries []models.BatteryTelemetry `json:"batteries"`
	Alerts    []models.BatteryAlert     `json:"alerts"`
}
