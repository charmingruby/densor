package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type EquipmentRepository interface {
	FindByName(name string) (entity.Equipment, error)
}
