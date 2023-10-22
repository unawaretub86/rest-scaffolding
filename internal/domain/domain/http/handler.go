package http

import (
	"github.com/unawaretub86/rest-scaffolding/internal/domain/domain/usecase"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/dependencies"
)

const suffixErr = "Error"

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
