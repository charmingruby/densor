package service

import "github.com/charmingruby/densor/internal/sensor/domain/repository"

func NewSensorService(
	sensorRepo repository.SensorRepository,
	sensorCategoryRepo repository.SensorCategoryRepository,
	sectorRepo repository.SectorRepository,
	equipmentRepo repository.EquipmentRepository) SensorService {
	return SensorService{
		sensorRepo:         sensorRepo,
		sensorCategoryRepo: sensorCategoryRepo,
		sectorRepo:         sectorRepo,
		equipmentRepo:      equipmentRepo,
	}
}

type SensorService struct {
	sensorRepo         repository.SensorRepository
	sensorCategoryRepo repository.SensorCategoryRepository
	sectorRepo         repository.SectorRepository
	equipmentRepo      repository.EquipmentRepository
}
