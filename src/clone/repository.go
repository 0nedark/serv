package clone

import (
	log "github.com/sirupsen/logrus"
	git "gopkg.in/src-d/go-git.v4"
)

// Repository to be cloned in the specified path
func Repository(url, path string) {
	logWithFields := log.WithFields(log.Fields{
		"url":  url,
		"path": path,
	})

	logWithFields.Info("Cloning repository")
	options := git.CloneOptions{URL: url}
	if _, err := git.PlainClone(path, false, &options); err == nil {
		logWithFields.Debug("Repository cloned")
	} else {
		logWithFields.WithError(err).Fatal("Failed to clone repository")
	}
}
