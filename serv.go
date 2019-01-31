package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"

	"github.com/0nedark/serv/src/command"
	"github.com/0nedark/serv/src/load"
	"github.com/0nedark/serv/src/verify"
)

func main() {
	app := cli.NewApp()
	app.Name = "serv"
	app.Usage = "Can be used to orchestrate commandup of services"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "specifies the serv.yml file",
			Value: "serv.yml",
		},
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "show all available logs",
		},
		cli.BoolFlag{
			Name:  "silent",
			Usage: "show only errors logs (ignored if verbose flag is provided)",
		},
	}

	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Unrecoverable error encountered")
	}
}

func action(c *cli.Context) error {
	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp: true,
		QuoteEmptyFields: true,
	})

	if c.GlobalBool("verbose") {
		log.SetLevel(log.DebugLevel)
	} else if c.GlobalBool("silent") {
		log.SetLevel(log.ErrorLevel)
	}

	serv := load.Config{}
	if err := load.GetConfig(c.GlobalString("file"), &serv); err != nil {
		return err
	}

	verify.Groups(serv.Order, serv.Groups)
	command.Groups(serv.Order, serv.Groups)

	return nil
}
