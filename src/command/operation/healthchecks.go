package operation

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/0nedark/serv/src/load"
)

type healthcheck struct {
	Path    string
	Command string
	Sleep   int
}

// Healthchecks of the service
func Healthchecks(service load.Service, lock *sync.WaitGroup) {
	healthchecksLock := &sync.WaitGroup{}
	defer healthchecksLock.Wait()

	for _, current := range service.Healthchecks {
		healthchecksLock.Add(1)
		timeout := time.Duration(current.Timeout) * time.Second
		hc := healthcheck{service.Path, current.Command, current.Sleep}
		go hc.start(timeout, healthchecksLock)
	}

	healthchecksLock.Wait()
	lock.Done()
}

func handleHealthcheckDone(logWithFields *log.Entry, success <-chan bool, timeout time.Duration) {
	select {
	case <-success:
		logWithFields.Debug("Healthcheck completed")
	case <-time.After(timeout):
		logWithFields.Fatal("Healthcheck timed out")
	}
}

func (hc healthcheck) start(timeout time.Duration, lock *sync.WaitGroup) {
	logWithFields := log.WithFields(log.Fields{
		"context": hc.Path,
		"command": hc.Command,
	})

	logWithFields.Info("Healthcheck starting")
	success := make(chan bool)
	go hc.loop(logWithFields, success)
	handleHealthcheckDone(logWithFields, success, timeout)
	lock.Done()
}

func (hc healthcheck) loop(logWithFields *log.Entry, success chan<- bool) {
	for {
		runCommand(
			handleCommand(hc.Path, hc.Command),
			handleHealthcheckError(logWithFields, success),
		)

		time.Sleep(time.Duration(hc.Sleep) * time.Second)
	}
}

func handleHealthcheckError(logWithFields *log.Entry, success chan<- bool) errorHandler {
	return func(err error) {
		if err != nil {
			logWithFields.WithError(err).Debug("Command failed")
		} else {
			success <- true
		}
	}
}
