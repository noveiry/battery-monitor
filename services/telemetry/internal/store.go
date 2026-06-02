package internal

import (
	"sync"

	"battery-monitor/pkg/models"
)

type Store struct {
	mu        sync.RWMutex
	batteries map[string]models.BatteryTelemetry
}

func NewStore() *Store {
	return &Store{
		batteries: make(map[string]models.BatteryTelemetry),
	}
}

func (s *Store) Upsert(b models.BatteryTelemetry) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.batteries[b.ID] = b
}

func (s *Store) Get(id string) (models.BatteryTelemetry, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	b, ok := s.batteries[id]

	return b, ok
}

func (s *Store) List() []models.BatteryTelemetry {
	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make(
		[]models.BatteryTelemetry,
		0,
		len(s.batteries),
	)

	for _, battery := range s.batteries {
		result = append(result, battery)
	}

	return result
}

type FleetStats struct {
	TotalBatteries int     `json:"total_batteries"`
	AvgCharge      float64 `json:"avg_charge"`
	AvgHealth      float64 `json:"avg_health"`
	AvgTemp        float64 `json:"avg_temperature"`
}

func (s *Store) Stats() FleetStats {

	s.mu.RLock()
	defer s.mu.RUnlock()

	if len(s.batteries) == 0 {
		return FleetStats{}
	}

	var chargeSum int
	var healthSum int
	var tempSum float64

	for _, b := range s.batteries {
		chargeSum += b.Charge
		healthSum += b.Health
		tempSum += b.Temperature
	}

	count := len(s.batteries)

	return FleetStats{
		TotalBatteries: count,
		AvgCharge:      float64(chargeSum) / float64(count),
		AvgHealth:      float64(healthSum) / float64(count),
		AvgTemp:        tempSum / float64(count),
	}
}
