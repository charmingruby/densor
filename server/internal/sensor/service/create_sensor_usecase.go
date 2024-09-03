package service

import (
	"log/slog"

	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/common/helper"
	"github.com/charmingruby/densor/internal/sensor/domain/entity"
	"github.com/charmingruby/densor/internal/sensor/domain/usecase"
)

const (
	createSensorUseCase = "Create Sensor Use Case"
)

func (s *SensorService) CreateSensorUseCase(in usecase.CreateSensorInputDTO) error {
	equipmentIsBlank := in.EquipmentName == ""
	sectorIsBlank := in.SectorName == ""

	if equipmentIsBlank && sectorIsBlank {
		return core.NewValidationErr("equipment name or sector name are required")
	}

	_, err := s.sensorRepo.FindByName(in.Name)
	if err == nil {
		return core.NewConflictErr("name")
	}

	category, err := s.sensorCategoryRepo.FindByName(in.SensorCategoryName)
	if err != nil {
		slog.Info(err.Error())
		return core.NewNotFoundErr("category")
	}

	var sector entity.Sector
	if !sectorIsBlank {
		sector, err = s.sectorRepo.FindByName(in.SectorName)
		if err != nil {
			slog.Info(err.Error())
			return core.NewNotFoundErr("sector")
		}
	}

	var equipment entity.Equipment
	if !equipmentIsBlank {
		equipment, err = s.equipmentRepo.FindByName(in.EquipmentName)
		if err != nil {
			slog.Info(err.Error())
			return core.NewNotFoundErr("equipment")
		}
	}

	sectorID := helper.Ternary(sectorIsBlank, "", sector.ID)
	equipmentID := helper.Ternary(equipmentIsBlank, "", equipment.ID)

	sensor := entity.NewSensor(entity.SensorInput{
		Name:             in.Name,
		Description:      in.Description,
		SensorCategoryID: category.ID,
		EquipmentID:      equipmentID,
		Observation:      in.Observation,
		SectorID:         sectorID,
	})

	if err := s.sensorRepo.Store(sensor); err != nil {
		slog.Error(err.Error())
		return core.NewInternalErr(createSensorUseCase)
	}

	return nil
}
