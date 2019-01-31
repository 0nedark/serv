package command

import (
	"sync"

	"github.com/0nedark/serv/src/command/operation"
	"github.com/0nedark/serv/src/load"
)

func commandGroup(service []load.Service) {
	run(operation.Start, service)
	run(operation.Healthchecks, service)
	run(operation.Postconditions, service)
}

func run(op operation.Handler, services []load.Service) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, service := range services {
		lock.Add(1)
		go op(service, lock)
	}
}
