package repository

import (
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

var gitClone = git.PlainClone

// Clone git repository in the specified path
func Clone(url, path string) (string, error) {
	log.WithFields(log.Fields{
		"url":  url,
		"path": path,
	}).Info("Cloning repository")

	options := git.CloneOptions{URL: url}
	_, err := gitClone(path, false, &options)
	if err == nil {
		return "Repository cloned", nil
	}

	return "Failed to clone repository", err
}
