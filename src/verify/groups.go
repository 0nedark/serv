package verify

import (
	"sync"

	"github.com/0nedark/serv/src/load"
)

// Groups verifies that repos in the group are present
func Groups(order []string, groups map[string][]load.Service) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, key := range order {
		verifyGroup(groups[key], lock)
	}
}
