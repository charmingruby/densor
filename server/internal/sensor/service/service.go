package service

import "github.com/charmingruby/densor/internal/sensor/domain/repository"

func NewSensorService(sensorRepo repository.SensorRepository) SensorService {
	return SensorService{
		sensorRepo: sensorRepo,
	}
}

type SensorService struct {
	sensorRepo repository.SensorRepository
}
