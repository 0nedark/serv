package verify

import (
	"github.com/0nedark/serv/src/load"
)

func getGroupsRepositories(groups load.Groups) []load.Repository {
	repositories := make([]load.Repository, 0)
	for _, service := range groups {
		repositories = append(repositories, getServiceRepositories(service)...)
	}

	return repositories
}

func getServiceRepositories(services []load.Service) []load.Repository {
	repositories := make([]load.Repository, 0)
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
