package read

import (
	"github.com/gin-gonic/gin"

	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
	yourHandler "github.com/unawaretub86/restScaffolding/internal/domain/yourDomain/http"
)

type read struct {
	container *dependencies.Container
}

func NewRead(container *dependencies.Container) *read {
	return &read{
		container: container,
	}
}

func (read *read) RegisterRoutes(basePath string, r *gin.Engine) {
	handler := yourHandler.NewHandler(read.container)

	r.GET("/ping", handler.Ping)
}
