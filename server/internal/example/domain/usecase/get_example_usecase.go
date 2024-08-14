package usecase

import (
	"github.com/charmingruby/densor/internal/common/core"
	"github.com/charmingruby/densor/internal/example/domain/entity"
)

func (s *ExampleUseCaseRegistry) GetExampleUseCase(id string) (*entity.Example, error) {
	example, err := s.exampleRepo.FindByID(id)
	if err != nil {
		return nil, core.NewNotFoundErr("example")
	}

	return example, nil
}
