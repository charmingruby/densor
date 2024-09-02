package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SensorRepository interface {
	FindByName(name string) (entity.Sensor, error)
	FindMany() ([]entity.Sensor, error)
	Store(sensor entity.Sensor) error
}
