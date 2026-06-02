package internal

import (
	"fmt"
	"math/rand"
	"time"

	"battery-monitor/pkg/models"
)

type Generator struct {
	id          string
	charge      int
	temperature float64
	health      int
	charging    bool
}

func NewGenerator(id string) *Generator {
	return &Generator{
		id:          id,
		charge:      rand.Intn(80) + 20,
		temperature: 25 + rand.Float64()*10,
		health:      rand.Intn(20) + 80,
		charging:    rand.Intn(2) == 0,
	}
}

func NewRandomID() string {
	return fmt.Sprintf("battery-%04d", rand.Intn(9999))
}

func (g *Generator) Next() models.BatteryTelemetry {

	if g.charging {
		g.charge += rand.Intn(4)

		if g.charge >= 100 {
			g.charge = 100
			g.charging = false
		}
	} else {

		g.charge -= rand.Intn(3)

		if g.charge <= 5 {
			g.charging = true
		}
	}

	if g.charge < 0 {
		g.charge = 0
	}

	g.temperature += (rand.Float64() * 4) - 2

	if g.temperature < 15 {
		g.temperature = 15
	}

	if g.temperature > 70 {
		g.temperature = 70
	}

	if rand.Intn(1000) < 5 {
		g.temperature = 55 + rand.Float64()*10
	}

	if rand.Intn(5000) < 3 {
		g.health -= 1
	}

	if g.health < 50 {
		g.health = 50
	}

	status := models.StatusDischarging

	if g.charging {
		status = models.StatusCharging
	}

	return models.BatteryTelemetry{
		ID:          g.id,
		Charge:      g.charge,
		Temperature: g.temperature,
		Health:      g.health,
		Status:      status,
		Timestamp:   time.Now().UTC(),
	}
}
