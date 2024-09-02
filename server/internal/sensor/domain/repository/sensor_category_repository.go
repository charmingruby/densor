package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SensorCategoryRepository interface {
	FindByName(name string) (entity.SensorCategory, error)
}
