package app

import (
	"io/ioutil"

	"github.com/0nedark/serv/src/command"
	"github.com/0nedark/serv/src/verify"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

func (a Application) action(c *cli.Context) error {
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
	config, err := a.loadConfig(file, ioutil.ReadFile)
	if err == nil {
		verify.Groups(config.Order, config.Groups)
		command.Groups(config.Order, config.Groups)
	}

	return err
}
