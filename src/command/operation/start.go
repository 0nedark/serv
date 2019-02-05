package operation

import (
	"sync"

	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Start the service
func Start(service load.Service, lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"context": service.Path,
		"command": service.Command,
	})

	logWithFields.Info("Service starting")
	runCommand(
		handleCommand(service.Path, service.Command),
		handleGenericError(logWithFields, "Service"),
	)

	lock.Done()
}
