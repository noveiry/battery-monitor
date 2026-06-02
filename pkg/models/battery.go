package models

import "time"

type BatteryStatus string

const (
	StatusCharging    BatteryStatus = "charging"
	StatusDischarging BatteryStatus = "discharging"
	StatusIdle        BatteryStatus = "idle"
)

type BatteryTelemetry struct {
	ID          string        `json:"id"`
	Charge      int           `json:"charge"`
	Temperature float64       `json:"temperature"`
	Health      int           `json:"health"`
	Status      BatteryStatus `json:"status"`
	Timestamp   time.Time     `json:"timestamp"`
}

type BatteryAlert struct {
	BatteryID string    `json:"battery_id"`
	Severity  string    `json:"severity"`
	Message   string    `json:"message"`
	Time      time.Time `json:"time"`
}