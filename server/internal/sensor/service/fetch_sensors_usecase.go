package service

import (
	"log/slog"

	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
)

const (
	fetchSensorsUseCase = "Fetch Sensors Use Case"
)

func (s *SensorService) FetchSensorsUseCase() ([]entity.Sensor, error) {
	sensors, err := s.sensorRepo.FindMany()
	if err != nil {
		slog.Error(err.Error())
		return nil, core.NewInternalErr(fetchSensorsUseCase)
	}

	return sensors, nil
}
