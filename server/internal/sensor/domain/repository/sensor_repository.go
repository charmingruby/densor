package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SensorRepository interface {
	Store(sensor entity.Sensor) error
	FindMany() ([]entity.Sensor, error)
}
