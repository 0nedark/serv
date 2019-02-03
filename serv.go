package main

import (
	"os"

	"github.com/0nedark/serv/app"
	"github.com/0nedark/serv/src/command"
	"github.com/0nedark/serv/src/load"
	"github.com/0nedark/serv/src/verify"
	log "github.com/sirupsen/logrus"
)

func main() {
	application := app.NewApplication(load.NewConfig, verify.Each, command.Groups)
	if err := application.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Unrecoverable error encountered")
	}
}
