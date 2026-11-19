package main

import (
	"pos/internal/cli"
	"pos/internal/container"
)

func main() {
	app := container.NewContainer()

	console := cli.NewCLI(app.Facade)
	console.Run()
}
