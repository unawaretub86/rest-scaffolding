// Package dependencies provider dependencies
package dependencies

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/configuration/database"
)

func StartMockDependencies() (*Container, sqlmock.Sqlmock) {

	mockdb, mock, _ := sqlmock.New()

	db := database.NewMockService(mockdb)

	// done
	return &Container{
		database:   db,
	}, mock
}
