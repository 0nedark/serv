package command

import (
	"github.com/0nedark/serv/src/load"
	log "github.com/sirupsen/logrus"
)

// GroupsFunc defines the signature of groups functions
type GroupsFunc = func([]string, load.Groups)

// Groups to be started
func Groups(order []string, groups load.Groups) {
	for _, key := range order {
		log.WithField("name", key).Info("Running group")
		commandGroup(groups[key])
	}
}
