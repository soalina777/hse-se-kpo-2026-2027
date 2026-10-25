package main

import (
	"vyshkat/container"
	"vyshkat/domain"
)

func main() {
	serviceCenter := domain.DefaultServiceCenter{}

	app := container.NewContainer(serviceCenter)
	app.Console.Run()
}
