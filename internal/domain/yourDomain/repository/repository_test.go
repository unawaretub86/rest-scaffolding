package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	database "github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/repository/database/mocks"
	"github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/repository"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
)

func Test_NewRepository(t *testing.T) {
	ass := assert.New(t)

	container := &dependencies.Container{}

	repo := repository.NewRepository(container)

	ass.NotNil(repo)
}

func Test_NewMockRepository(t *testing.T) {
	ass := assert.New(t)

	databaseRepository := database.NewMockDatabase()

	repo := repository.NewMockRepository(databaseRepository)

	ass.NotNil(repo)
}
