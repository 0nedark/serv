package operation

import (
	"os/exec"
	"sync"

	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Start the service
func Start(service load.Service, lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"command": service.Command.Name,
		"args":    service.Command.Args,
	})

	logWithFields.Info("Service starting")

	command := exec.Command(service.Command.Name, service.Command.Args...)
	command.Dir = buildPath(service.Path)
	if stdout, err := command.Output(); err != nil {
		logWithFields.WithError(err).Fatal("Service start failed")
	} else {
		logWithFields.Debug("Service started")
		log.Debug("Command output:\n" + string(stdout))
	}

	lock.Done()
}
