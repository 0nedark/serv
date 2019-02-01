package operation

import (
	"sync"

	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

type postcondition struct {
	Path    string
	Command string
}

// Postconditions of the service
func Postconditions(service load.Service, lock *sync.WaitGroup) {
	postconditionsLock := &sync.WaitGroup{}
	defer postconditionsLock.Wait()

	for _, current := range service.Postconditions {
		postconditionsLock.Add(1)
		pc := postcondition{service.Path, current.Command}
		go pc.start(postconditionsLock)
	}

	postconditionsLock.Wait()
	lock.Done()
}

func (pc postcondition) start(lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"context": pc.Path,
		"command": pc.Command,
	})

	logWithFields.Info("Postcondition starting")
	output := runCommand(
		handleCommand(pc.Path, pc.Command),
		handleGenericError(logWithFields, "Postcondition"),
	)

	log.Debugf("Command output:\n%s", output)
	lock.Done()
}
