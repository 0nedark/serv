package app

import cli "gopkg.in/urfave/cli.v1"

// Configure the cli application
func Configure() *cli.App {
	app := cli.NewApp()
	app.Name = "serv"
	app.Usage = "Orchestrate the startup of your services"
	app.Version = "1.0.0"
	app.Action = action
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

	return app
}
