package app

import (
	"github.com/0nedark/serv/src/command"
	"github.com/0nedark/serv/src/load"
	"github.com/0nedark/serv/src/verify"
	cli "gopkg.in/urfave/cli.v1"
)

// Application defines cli application injections
type Application struct {
	instance      *cli.App
	loadConfig    load.ConfigFunc
	verifyGroups  verify.GroupsFunc
	commandGroups command.GroupsFunc
}

// Run the cli application
func (a Application) Run(args []string) error {
	return a.instance.Run(args)
}

// NewApplication constructs instance of application
func NewApplication(
	loadConfig load.ConfigFunc,
	verifyGroups verify.GroupsFunc,
	commandGroups command.GroupsFunc,
) Application {
	application := Application{cli.NewApp(), loadConfig, verifyGroups, commandGroups}
	application.configure()

	return application
}
