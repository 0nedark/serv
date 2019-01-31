package operation

import (
	"os/exec"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/drupsys/serv/src/load"
)

// Healthchecks of the service
func Healthchecks(service load.Service, lock *sync.WaitGroup) {
	healthchecksLock := &sync.WaitGroup{}
	defer healthchecksLock.Wait()

	for _, hc := range service.Healthchecks {
		healthchecksLock.Add(1)
		go healthcheckInit(hc, healthchecksLock)
	}

	healthchecksLock.Wait()
	lock.Done()
}

func healthcheckInit(hc load.Healthcheck, lock *sync.WaitGroup) {
	timeout := time.Duration(hc.Timeout) * time.Second
	result := make(chan bool)
	go healthcheckLoop(result, hc)

	logWithFields := log.WithFields(log.Fields{
		"command": hc.Name,
		"args":    hc.Args,
	})

	select {
	case <-result:
		logWithFields.Info("Healthcheck competed")
	case <-time.After(timeout):
		logWithFields.Fatal("Healthcheck timed out")
	}

	lock.Done()
}

func healthcheckLoop(result chan bool, hc load.Healthcheck) {
	for {
		healthcheck(result, hc)
		time.Sleep(time.Duration(hc.Sleep) * time.Second)
	}
}

func healthcheck(result chan bool, hc load.Healthcheck) {
	logWithFields := log.WithFields(log.Fields{
		"command": hc.Name,
		"args":    hc.Args,
	})

	command := exec.Command(hc.Name, hc.Args...)
	if stdout, err := command.Output(); err != nil {
		logWithFields.WithError(err).Debug("Healthcheck failed")
	} else {
		logWithFields.Debug(string(stdout))
		result <- true
	}
}
