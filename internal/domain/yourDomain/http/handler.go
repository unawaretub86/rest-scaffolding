package http

import (
	"github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/usecase"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
)

type (
	Handler struct {
		UseCase usecase.UseCase
	}
)

func NewHandler(container *dependencies.Container) *Handler {
	return &Handler{
		UseCase: usecase.NewUse(container),
	}
}
