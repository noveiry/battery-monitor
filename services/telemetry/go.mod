module telemetry

go 1.24

require (
	github.com/nats-io/nats.go v1.46.0
)

replace battery-fleet-monitor/pkg/models => ../../pkg/models
replace battery-fleet-monitor/pkg/events => ../../pkg/events
replace battery-fleet-monitor/pkg/natsx => ../../pkg/natsx
replace battery-fleet-monitor/pkg/logger => ../../pkg/logger