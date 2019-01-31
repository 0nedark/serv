package verify

import (
	"sync"

	"github.com/drupsys/serv/src/load"
)

func verifyGroup(group []load.Service) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, service := range group {
		isRepository(service.Repository, lock)
	}
}
