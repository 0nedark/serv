package verify

import (
	"sync"

	"github.com/0nedark/serv/src/load"
)

func verifyRepositories(repositories []load.Repository) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, repository := range repositories {
		lock.Add(1)
		go verifyRepository(repository, lock)
	}
}
