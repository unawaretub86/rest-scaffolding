package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/usecase"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
)

func Test_NewUse(t *testing.T) {
	ass := assert.New(t)

	uc := usecase.NewUse(&dependencies.Container{})

	ass.NotNil(uc)
}

func Test_NewUseMockUser(t *testing.T) {
	ass := assert.New(t)

	usecaseMock := usecase.NewUse(&dependencies.Container{})

	uc := usecase.NewMockUse(usecaseMock)

	ass.NotNil(uc)
}
