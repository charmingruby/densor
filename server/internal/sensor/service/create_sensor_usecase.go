package service

import (
	"log/slog"

	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/charmingruby/densor/internal/sensor/domain/usecase"
)

const (
	createSensorUseCase = "Create Sensor Use Case"
)

func (s *SensorService) CreateSensorUseCase(in usecase.CreateSensorInputDTO) error {
	sensor := entity.NewSensor(entity.SensorInput{
		Name:             in.Name,
		Description:      in.Description,
		SensorCategoryID: in.SensorCategoryID,
		EquipmentID:      in.EquipmentID,
		Observation:      in.Observation,
		SectorID:         in.SectorID,
	})

	if err := s.sensorRepo.Store(sensor); err != nil {
		slog.Error(err.Error())
		return core.NewInternalErr(createSensorUseCase)
	}

	return nil
}
