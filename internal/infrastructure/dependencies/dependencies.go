package dependencies

import (
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/configuration"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/configuration/database"
)

type Container struct {
	database database.Database
}

func StartDependencies() *Container {
	config, _ := configuration.LoadConfiguration(".")

	db, err := database.ConnectDB(config.DB_URL)
	if err != nil {
		panic(err.Error())
	}

	container := &Container{
		database: db,
	}

	return container
}

func (container *Container) Database() database.Database {
	return container.database
}
