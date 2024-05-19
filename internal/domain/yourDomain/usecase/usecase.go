package usecase

import (
	"github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/repository"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
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

func NewMockUse(repositoryUser repository.Repo) UseCase {
	return &useCase{
		repo: repositoryUser,
	}
}
