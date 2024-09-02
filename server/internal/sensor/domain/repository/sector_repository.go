package repository

import "github.com/charmingruby/densor/internal/sensor/domain/entity"

type SectorRepository interface {
	FindByName(name string) (entity.Sector, error)
}
