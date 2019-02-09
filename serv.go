package main

import (
	"os"

	"github.com/0nedark/serv/app"
	log "github.com/sirupsen/logrus"
)

var runAppWith = app.Configure().Run
var fatal = log.Fatal

func main() {
	if err := runAppWith(os.Args); err != nil {
		fatal(err.Error())
	}
}
