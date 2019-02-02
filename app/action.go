package app

import (
	"io/ioutil"

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
		a.verifyGroups(config.Groups)
		a.commandGroups(config.Order, config.Groups)
	}

	return err
}
