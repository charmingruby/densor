package usecase

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SensorService interface {
	CreateSensorUseCase(dto CreateSensorInputDTO) error
	FetchSensorsUseCase() ([]entity.Sensor, error)
}

type CreateSensorInputDTO struct {
	Name             string
	Description      string
	SensorCategoryID string
	EquipmentID      string
	Observation      string
	SectorID         string
}
