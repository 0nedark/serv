package main

import (
	"os"

	"github.com/0nedark/serv/app"
	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

func main() {
	application := app.NewApplication(load.NewConfig)
	if err := application.Run(os.Args); err != nil {
		log.WithError(err).Fatal("Unrecoverable error encountered")
	}
}
