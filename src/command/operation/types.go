package operation

import (
	"os"
	"path"
	"sync"

	"github.com/drupsys/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Handler defines the signature of operation
type Handler = func(load.Service, *sync.WaitGroup)

func buildPath(dir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("Unable to get current working directory")
	}

	return path.Join(cwd, dir)
}
