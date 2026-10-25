package container

import (
	"vyshkat/cli"
	"vyshkat/domain"
	"vyshkat/service"
)

type Container struct {
	Console *cli.Console
}

func NewContainer(serviceCenter domain.ServiceCenter) *Container {
	parkService := service.NewParkService(serviceCenter)

	return &Container{
		Console: cli.NewConsole(parkService),
	}
}
