package app

import cli "gopkg.in/urfave/cli.v1"

func (a Application) configure() {
	a.instance.Name = "serv"
	a.instance.Usage = "Orchestrate the startup of your services"
	a.instance.Version = "1.0.0"
	a.instance.Action = a.action
	a.instance.Flags = []cli.Flag{
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
}
