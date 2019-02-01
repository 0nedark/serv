package operation

import (
	"os/exec"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/0nedark/serv/src/load"
)

// Healthchecks of the service
func Healthchecks(service load.Service, lock *sync.WaitGroup) {
	healthchecksLock := &sync.WaitGroup{}
	defer healthchecksLock.Wait()

	for _, hc := range service.Healthchecks {
		healthchecksLock.Add(1)
		go healthcheckInit(service.Path, hc, healthchecksLock)
	}

	healthchecksLock.Wait()
	lock.Done()
}

func healthcheckInit(path string, hc load.Healthcheck, lock *sync.WaitGroup) {
	logWithFields := log.WithField("command", hc.Command)
	logWithFields.Info("Healthcheck started")

	timeout := time.Duration(hc.Timeout) * time.Second
	result := make(chan bool)
	go healthcheckLoop(result, path, hc)

	select {
	case <-result:
		logWithFields.Debug("Healthcheck completed")
	case <-time.After(timeout):
		logWithFields.Fatal("Healthcheck timed out")
	}

	lock.Done()
}

func healthcheckLoop(result chan bool, path string, hc load.Healthcheck) {
	for {
		healthcheck(result, path, hc)
		time.Sleep(time.Duration(hc.Sleep) * time.Second)
	}
}

func healthcheck(result chan bool, path string, hc load.Healthcheck) {
	logWithFields := log.WithField("command", hc.Command)

	command := exec.Command("sh", "-c", hc.Command)
	command.Dir = buildPath(path)
	stdout, err := command.Output()

	if err != nil {
		logWithFields.WithError(err).Debug("Healthcheck failed")
	} else {
		result <- true
	}

	log.Debug("Command output:\n" + string(stdout))
}
