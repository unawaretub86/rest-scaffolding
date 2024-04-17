package repository

import (
	database "github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/repository/database"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
)

type (
	Repo interface {
	}

	repository struct {
		database database.Database
	}
)

func NewRepository(container *dependencies.Container) Repo {
	return &repository{
		database: database.NewDatabase(container),
	}
}