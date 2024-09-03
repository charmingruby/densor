package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SensorCategoryRepository interface {
	FindMany() ([]entity.SensorCategory, error)
	FindByName(name string) (entity.SensorCategory, error)
}
