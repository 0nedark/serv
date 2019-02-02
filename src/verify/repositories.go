package verify

import (
	"sync"

	"github.com/0nedark/serv/src/clone"
	"github.com/0nedark/serv/src/load"

	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

func selectRepositories(services []load.Service, repositories []load.Repository) []load.Repository {
	for _, service := range services {
		if !emptyRepository(service.Repository) {
			repositories = append(repositories, service.Repository)
		}
	}

	return repositories
}

func emptyRepository(repository load.Repository) bool {
	return repository == (load.Repository{})
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
