package main

import (
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/configuration/server"
	"github.com/unawaretub86/restScaffolding/internal/infrastructure/dependencies"
)

const basePath = ""

func main() {
	container := dependencies.StartDependencies()

	server.NewServer(basePath, container)
}