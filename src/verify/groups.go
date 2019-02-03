package verify

import (
	"sync"

	"github.com/0nedark/serv/src/load"
	"github.com/0nedark/serv/src/repository"

	log "github.com/sirupsen/logrus"
)

// EachFunc defines signature of Each exported function
type EachFunc = func(load.Groups)

// Each verifies that repositories in the group are present
func Each(groups load.Groups) {
	lock := &sync.WaitGroup{}
	defer lock.Wait()

	repositories := getGroupsRepositories(groups)
	for _, r := range repositories {
		lock.Add(1)
		go verify(r.URL, r.Path, lock)
	}
}

var open = repository.Open
var clone = repository.Clone

func verify(URL, Path string, lock *sync.WaitGroup) {
	if remote, err := open(Path); err != nil {
		clone(URL, Path)
	} else {
		log.WithFields(log.Fields{
			"url":  remote,
			"path": Path,
		}).Debug("Repository exists")
	}

	lock.Done()
}
