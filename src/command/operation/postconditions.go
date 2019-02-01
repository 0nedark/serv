package operation

import (
	"os/exec"
	"sync"

	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Postconditions of the service
func Postconditions(service load.Service, lock *sync.WaitGroup) {
	postconditionsLock := &sync.WaitGroup{}
	defer postconditionsLock.Wait()

	for _, postcondition := range service.Postconditions {
		postconditionsLock.Add(1)
		go runPostcondition(service.Path, postcondition, postconditionsLock)
	}

	postconditionsLock.Wait()
	lock.Done()
}

func runPostcondition(path string, pc load.Command, lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"command": pc.Name,
		"args":    pc.Args,
	})

	logWithFields.Info("Postcondition started")

	command := exec.Command(pc.Name, pc.Args...)
	command.Dir = buildPath(path)
	if stdout, err := command.Output(); err != nil {
		logWithFields.WithError(err).Fatal("Postcondition failed")
	} else {
		logWithFields.Debug("Postcondition completed")
		log.Debug("Command output:\n" + string(stdout))
	}

	lock.Done()
}
