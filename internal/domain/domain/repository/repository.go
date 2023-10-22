package repository

import (
	"github.com/unawaretub86/rest-scaffolding/internal/domain/domain/repository/database"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/dependencies"
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
