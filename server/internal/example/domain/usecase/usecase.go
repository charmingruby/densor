package usecase

import (
	"github.com/charmingruby/densor/internal/example/domain/dto"
	"github.com/charmingruby/densor/internal/example/domain/entity"
	"github.com/charmingruby/densor/internal/example/domain/repository"
)

type ExampleUseCase interface {
	CreateExampleUseCase(dto dto.CreateExampleUseCaseDTO) error
	GetExampleUseCase(id string) (*entity.Example, error)
}

func NewExampleUseCaseRegistry(exampleRepo repository.ExampleRepository) *ExampleUseCaseRegistry {
	return &ExampleUseCaseRegistry{
		exampleRepo: exampleRepo,
	}
}

type ExampleUseCaseRegistry struct {
	exampleRepo repository.ExampleRepository
}
