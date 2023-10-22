package main

import (
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/configuration/server"
	"github.com/unawaretub86/rest-scaffolding/internal/infrastructure/dependencies"
)

const basePath = ""

func main() {
	container := dependencies.StartDependencies()

	server.NewServer(basePath, container)
}
