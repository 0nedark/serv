package app

import (
	"github.com/0nedark/serv/src/load"
	cli "gopkg.in/urfave/cli.v1"
)

// GroupsFunc defines the signature of groups functions
type GroupsFunc = func([]string, load.Groups)

// Application defines cli application injections
type Application struct {
	instance      *cli.App
	loadConfig    load.ConfigFunc
	verifyGroups  GroupsFunc
	commandGroups GroupsFunc
}

// Run the cli application
func (a Application) Run(args []string) error {
	return a.instance.Run(args)
}

// NewApplication constructs instance of application
func NewApplication(loadConfig load.ConfigFunc, verifyGroups, commandGroups GroupsFunc) Application {
	application := Application{cli.NewApp(), loadConfig, verifyGroups, commandGroups}
	application.configure()

	return application
}
