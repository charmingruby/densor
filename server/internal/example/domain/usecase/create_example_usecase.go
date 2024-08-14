package usecase

import (
	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/example/domain/dto"
	"github.com/charmingruby/densor/internal/example/domain/entity"
)

func (s *ExampleUseCaseRegistry) CreateExampleUseCase(dto dto.CreateExampleUseCaseDTO) error {
	example, err := entity.NewExample(dto.Name)
	if err != nil {
		return err
	}

	if err := s.exampleRepo.Store(example); err != nil {
		return core.NewInternalErr("create example store")
	}

	return nil
}
