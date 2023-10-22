package usecase

import (
	"github.com/unawaretub86/rest-scaffolding/internal/domain/domain/repository"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/dependencies"
)

type (
	UseCase interface {
	}

	useCase struct {
		repo repository.Repo
	}
)

func NewUse(container *dependencies.Container) UseCase {
	return &useCase{
		repo: repository.NewRepository(container),
	}
}
