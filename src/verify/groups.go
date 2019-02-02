package verify

import (
	"github.com/0nedark/serv/src/load"
)

// Groups verifies that repositories in the group are present
func Groups(order []string, groups load.Groups) {
	repositories := make([]load.Repository, 0)
	for _, key := range order {
		repositories = selectRepositories(groups[key], repositories)
	}

	verifyRepositories(repositories)
}
