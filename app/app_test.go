package app

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/0nedark/serv/src/load"

	log "github.com/sirupsen/logrus"
)

func getConfigStub(string, load.ReadFileFunc) (load.Config, error) {
	return load.Config{}, nil
}

func eachStub(groups load.Groups)                          {}
func commandGroupsStub(order []string, groups load.Groups) {}

func TestAppPackage(t *testing.T) {
	Convey("package app", t, func() {
		application := NewApplication(
			getConfigStub,
			eachStub,
			commandGroupsStub,
		)

		Convey("application.Run", func() {
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
	})
}
