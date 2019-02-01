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

type commandHandler = func() (string, error)
type errorHandler = func(error)

func handleCommand(path, command string) commandHandler {
	return func() (string, error) {
		exec := exec.Command("sh", "-c", command)
		exec.Dir = buildPath(path)
		stdout, err := exec.Output()

		return string(stdout), err
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

func runCommand(handleCommand commandHandler, handleError errorHandler) string {
	output, err := handleCommand()
	handleError(err)

	return output
}
