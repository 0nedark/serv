package verify

import (
	"sync"

	"github.com/0nedark/serv/src/load"
)

// Groups verifies that repositories in the group are present
func Groups(order []string, groups load.Groups) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	for _, key := range order {
		verifyGroup(groups[key], lock)
	}
}
