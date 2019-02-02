package app

import (
	"github.com/0nedark/serv/src/load"
	cli "gopkg.in/urfave/cli.v1"
)

type Application struct {
	instance   *cli.App
	loadConfig load.ConfigFunc
}

func NewApplication(loadConfig load.ConfigFunc) Application {
	application := Application{cli.NewApp(), loadConfig}
	application.configure()

	return application
}
