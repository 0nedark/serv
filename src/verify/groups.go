package verify

import "github.com/drupsys/serv/src/load"

// Groups verifies that repos in the group are present
func Groups(order []string, groups map[string][]load.Service) {
	for _, key := range order {
		verifyGroup(groups[key])
	}
}
