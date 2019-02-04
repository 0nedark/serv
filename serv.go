package main

import (
	"os"

	"github.com/0nedark/serv/app"
	log "github.com/sirupsen/logrus"
)

var run = app.Configure().Run
var fatal = log.Fatal

func main() {
	if err := run(os.Args); err != nil {
		fatal(err.Error())
	}
}
