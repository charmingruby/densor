package entity

import (
	"time"

	"github.com/google/uuid"
)

type SensorInput struct {
	Name             string
	Description      string
	SensorCategoryID string
	EquipmentID      string
	Observation      string
	SectorID         string
}

func NewSensor(in SensorInput) Sensor {
	return Sensor{
		ID:               uuid.New().String(),
		Name:             in.Name,
		Description:      in.Description,
		SensorCategoryID: in.SensorCategoryID,
		EquipmentID:      in.EquipmentID,
		Status:           "OK",
		Observation:      in.Observation,
		SectorID:         in.SectorID,
		CreatedAt:        time.Now(),
	}
}

type Sensor struct {
	ID               string
	Name             string
	Description      string
	SensorCategoryID string
	EquipmentID      string
	Status           string
	Observation      string
	SectorID         string
	CreatedAt        time.Time
}
