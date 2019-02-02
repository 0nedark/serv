package verify

import (
	"github.com/0nedark/serv/src/load"
)

// GroupsFunc defines the signature of groups functions
type GroupsFunc = func(load.Groups)

// Groups verifies that repositories in the group are present
func Groups(groups load.Groups) {
	repositories := make([]load.Repository, 0)
	for _, service := range groups {
		repositories = selectRepositories(service, repositories)
	}

	verifyRepositories(repositories)
}
