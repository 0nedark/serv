package main

import (
	"log"
	"os"

	"github.com/drupsys/serv/src/command"
	"github.com/drupsys/serv/src/load"
	"github.com/drupsys/serv/src/verify"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "serv"
	app.Usage = "Can be used to orchestrate commandup of services"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "Specifies the serv.yml file",
			Value: "serv.yml",
		},
	}

	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	serv := load.Config{}
	load.GetConfig(c.GlobalString("file"), &serv)
	verify.Groups(serv.Order, serv.Groups)
	command.Groups(serv.Order, serv.Groups)

	return nil
}
