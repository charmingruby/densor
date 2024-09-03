package service

import (
	"log/slog"

	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
)

const (
	fetchSensorCategoriesUseCase = "Fetch Sensor Categories Use Case"
)

func (s *SensorService) FetchSensorCategoriesUseCase() ([]entity.SensorCategory, error) {
	categories, err := s.sensorCategoryRepo.FindMany()
	if err != nil {
		slog.Error(err.Error())
		return nil, core.NewInternalErr(fetchSensorCategoriesUseCase)
	}

	return categories, nil
}
