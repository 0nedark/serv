package verify

import (
	"sync"

	"github.com/0nedark/serv/src/load"
)

func verifyGroup(group []load.Service) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, service := range group {
		isRepository(service.Repository, lock)
	}
}
