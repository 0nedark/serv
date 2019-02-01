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

func runPostcondition(path string, pc load.Postcondition, lock *sync.WaitGroup) {
	logWithFields := log.WithField("command", pc.Command)
	logWithFields.Info("Postcondition started")

	command := exec.Command("sh", "-c", pc.Command)
	command.Dir = buildPath(path)
	stdout, err := command.Output()
	if err != nil {
		logWithFields.WithError(err).Fatal("Postcondition failed")
	} else {
		logWithFields.Debug("Postcondition completed")
	}

	log.Debug("Command output:\n" + string(stdout))

	lock.Done()
}
