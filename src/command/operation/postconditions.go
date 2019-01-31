package operation

import (
	"os/exec"
	"sync"

	"github.com/drupsys/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// Postconditions of the service
func Postconditions(service load.Service, lock *sync.WaitGroup) {
	postconditionsLock := &sync.WaitGroup{}
	defer postconditionsLock.Wait()

	for _, postcondition := range service.Postconditions {
		postconditionsLock.Add(1)
		go runPostcondition(postcondition, postconditionsLock)
	}

	postconditionsLock.Wait()
	lock.Done()
}

func runPostcondition(pc load.Command, lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"command": pc.Name,
		"args":    pc.Args,
	})

	command := exec.Command(pc.Name, pc.Args...)
	if stdout, err := command.Output(); err != nil {
		logWithFields.WithError(err).Fatal("Postcondition failed")
	} else {
		logWithFields.Info("Postcondition completed")
		logWithFields.Debug(string(stdout))
	}

	lock.Done()
}
