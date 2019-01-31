package command

import (
	"fmt"

	"github.com/drupsys/serv/src/load"
)

// Groups to be started
func Groups(order []string, groups map[string][]load.Service) {
	for _, key := range order {
		fmt.Printf("\n\t:%s:\n\n", key)

		commandGroup(groups[key])
	}

	fmt.Println()
}
