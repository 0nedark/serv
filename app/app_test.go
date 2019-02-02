package app

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/0nedark/serv/src/load"

	log "github.com/sirupsen/logrus"
)

func getConfigMock(string, load.ReadFileFunc) (load.Config, error) {
	return load.Config{}, nil
}

func TestCommandIsStarting(t *testing.T) {
	Convey("Flags are handled correctly", t, func() {
		application := NewApplication(getConfigMock)

		Convey("should set log level to debug with --verbose", func() {
			application.Run([]string{"pwd", "--verbose"})
			So(log.GetLevel(), ShouldEqual, log.DebugLevel)
		})

		Convey("should set log level to fatal with --silent", func() {
			application.Run([]string{"pwd", "--silent"})
			So(log.GetLevel(), ShouldEqual, log.FatalLevel)
		})

		Convey("should ignore --silent flag if --verbose is set", func() {
			application.Run([]string{"pwd", "--silent", "--verbose"})
			So(log.GetLevel(), ShouldEqual, log.DebugLevel)
		})
	})
}
