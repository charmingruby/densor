package mapper

import (
	"time"

	"github.com/charmingruby/densor/internal/sensor/domain/entity"
)

func DomainSensorToPostgres(sensor entity.Sensor) PostgresSensor {
	return PostgresSensor{
		ID:               sensor.ID,
		Name:             sensor.Name,
		Description:      sensor.Description,
		SensorCategoryID: sensor.SensorCategoryID,
		EquipmentID:      stringToPointer(sensor.EquipmentID),
		Status:           sensor.Status,
		Observation:      sensor.Observation,
		SectorID:         stringToPointer(sensor.SectorID),
		CreatedAt:        sensor.CreatedAt,
	}
}

func PostgresSensorToDomain(pgSensor PostgresSensor) entity.Sensor {
	return entity.Sensor{
		ID:               pgSensor.ID,
		Name:             pgSensor.Name,
		Description:      pgSensor.Description,
		SensorCategoryID: pgSensor.SensorCategoryID,
		EquipmentID:      pointerToString(pgSensor.EquipmentID),
		Status:           pgSensor.Status,
		Observation:      pgSensor.Observation,
		SectorID:         pointerToString(pgSensor.SectorID),
		CreatedAt:        pgSensor.CreatedAt,
	}
}

type PostgresSensor struct {
	ID               string    `json:"id" db:"id"`
	Name             string    `json:"name" db:"name"`
	Description      string    `json:"description" db:"description"`
	SensorCategoryID string    `json:"sensor_category_id" db:"sensor_category_id"`
	Status           string    `json:"status" db:"status"`
	Observation      string    `json:"observation" db:"observation"`
	EquipmentID      *string   `json:"equipment_id" db:"equipment_id"`
	SectorID         *string   `json:"sector_id" db:"sector_id"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
}
