package database

import (
	database "github.com/unawaretub86/rest-scaffolding/internal/infrastructure/configuration/database"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/dependencies"
)

type (
	Database interface {
	}

	databaseDomain struct {
		db database.Database
	}
)

func NewDatabase(container *dependencies.Container) Database {
	return &databaseDomain{
		db: container.Database(),
	}
}
