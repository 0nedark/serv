package main

import (
	"os"

	"github.com/0nedark/serv/app"
	log "github.com/sirupsen/logrus"
)

func main() {
	if err := app.Configure().Run(os.Args); err != nil {
		log.WithError(err).Fatal("Unrecoverable error encountered")
	}
}
