package internal

import (
	"battery-monitor/pkg/events"
	"battery-monitor/pkg/models"
	"battery-monitor/pkg/natsx"

	"github.com/nats-io/nats.go"
)

type Publisher struct {
	conn *nats.Conn
}

func NewPublisher(conn *nats.Conn) *Publisher {
	return &Publisher{
		conn: conn,
	}
}

func (p *Publisher) Publish(
	event models.BatteryTelemetry,
) error {

	return natsx.Publish(
		p.conn,
		events.BatteryTelemetrySubject,
		event,
	)
}
