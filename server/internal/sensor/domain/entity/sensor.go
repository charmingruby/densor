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
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Description      string    `json:"description"`
	SensorCategoryID string    `json:"sensor_category_id"`
	EquipmentID      string    `json:"equipment_id"`
	Status           string    `json:"status"`
	Observation      string    `json:"observation"`
	SectorID         string    `json:"sector_id"`
	CreatedAt        time.Time `json:"created_at"`
}
