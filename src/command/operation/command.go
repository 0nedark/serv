package operation

import (
	"os"
	"os/exec"
	"path"

	log "github.com/sirupsen/logrus"
)

func buildPath(dir string) string {
	cwd, err := os.Getwd()
	if err != nil {
		log.WithError(err).Fatal("Unable to get current working directory")
	}

	return path.Join(cwd, dir)
}

type commandHandler = func() error
type errorHandler = func(error)

func handleCommand(path, command string) commandHandler {
	return func() error {
		exec := exec.Command("sh", "-c", command)
		exec.Dir = buildPath(path)
		if log.IsLevelEnabled(log.DebugLevel) {
			exec.Stderr = os.Stderr
			exec.Stdout = os.Stdout
		}

		return exec.Run()
	}
}

func handleGenericError(logWithFields *log.Entry, prefix string) errorHandler {
	return func(err error) {
		if err != nil {
			logWithFields.WithError(err).Fatalf("%s failed", prefix)
		} else {
			logWithFields.Debugf("%s started", prefix)
		}
	}
}

func runCommand(handleCommand commandHandler, handleError errorHandler) {
	handleError(handleCommand())
}
