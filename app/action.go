package app

import (
	"github.com/0nedark/serv/src/command"
	"github.com/0nedark/serv/src/load"
	"github.com/0nedark/serv/src/verify"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

var loadConfig = load.Config
var verifyEach = verify.Each
var commandGroups = command.Groups

func action(c *cli.Context) error {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
		QuoteEmptyFields: true,
	})

	if c.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	} else if c.GlobalBool("silent") {
		log.SetLevel(log.FatalLevel)
	}

	file := c.GlobalString("file")
	config, err := loadConfig(file)
	if err == nil {
		verifyEach(config.Groups)
		commandGroups(config.Order, config.Groups)
	}

	return err
}
