package repository

import git "gopkg.in/src-d/go-git.v4"

// Open will try to open and return remote of local repository
func Open(path string) (string, error) {
	repository, err := git.PlainOpen(path)
	if err == nil {
		remote, _ := repository.Remote("origin")
		return remote.Config().URLs[0], nil
	}

	return "", err
}
