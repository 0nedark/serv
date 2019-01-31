package verify

import (
	"sync"

	"github.com/0nedark/serv/src/clone"
	"github.com/0nedark/serv/src/load"

	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

func isRepository(repository load.Repository, lock *sync.WaitGroup) {
	if repository != (load.Repository{}) {
		lock.Add(1)
		go verifyRepository(repository, lock)
	}
}

func verifyRepository(repository load.Repository, lock *sync.WaitGroup) {
	if repo, err := git.PlainOpen(repository.Path); err != nil {
		clone.Repository(repository.URL, repository.Path)
	} else {
		remote, _ := repo.Remote("origin")
		log.WithFields(log.Fields{
			"url":  remote.Config().URLs[0],
			"path": repository.Path,
		}).Debug("Repository exists")
	}

	lock.Done()
}
