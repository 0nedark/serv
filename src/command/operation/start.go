package operation

import (
	"os/exec"
	"sync"

	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Start the service
func Start(service load.Service, lock *sync.WaitGroup) {
	logWithFields := log.WithField("command", service.Command)
	logWithFields.Info("Service starting")

	command := exec.Command("sh", "-c", service.Command)
	command.Dir = buildPath(service.Path)
	stdout, err := command.Output()

	if err != nil {
		logWithFields.WithError(err).Fatal("Service start failed")
	} else {
		logWithFields.Debug("Service started")
	}

	log.Debug("Command output:\n" + string(stdout))

	lock.Done()
}
