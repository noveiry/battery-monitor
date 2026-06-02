package internal

import (
	"sync"

	"battery-fleet-monitor/pkg/models"
)

type Store struct {
	mu     sync.RWMutex
	alerts []models.BatteryAlert
}

func NewStore() *Store {
	return &Store{
		alerts: make([]models.BatteryAlert, 0),
	}
}

func (s *Store) Add(
	alert models.BatteryAlert,
) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.alerts = append(s.alerts, alert)

	// keep only last 1000 alerts
	if len(s.alerts) > 1000 {
		s.alerts = s.alerts[len(s.alerts)-1000:]
	}
}

func (s *Store) All() []models.BatteryAlert {

	s.mu.RLock()
	defer s.mu.RUnlock()

	result := make([]models.BatteryAlert, len(s.alerts))
	copy(result, s.alerts)

	return result
}

func (s *Store) Critical() []models.BatteryAlert {

	s.mu.RLock()
	defer s.mu.RUnlock()

	var result []models.BatteryAlert

	for _, a := range s.alerts {
		if a.Severity == "critical" {
			result = append(result, a)
		}
	}

	return result
}
