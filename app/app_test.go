package app

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/0nedark/serv/src/load"

	log "github.com/sirupsen/logrus"
)

func eachStub(groups load.Groups)                          {}
func commandGroupsStub(order []string, groups load.Groups) {}
func getConfigStub(string, load.ReadFileFunc) (load.Config, error) {
	return load.Config{}, nil
}

func TestAppPackage(t *testing.T) {
	Convey("package app", t, func() {
		newConfig = getConfigStub
		verifyEach = eachStub
		commandGroups = commandGroupsStub

		Convey("Configure", func() {
			Convey("should set log level to debug with --verbose", func() {
				Configure().Run([]string{"pwd", "--verbose"})
				So(log.GetLevel(), ShouldEqual, log.DebugLevel)
			})

			Convey("should set log level to fatal with --silent", func() {
				Configure().Run([]string{"pwd", "--silent"})
				So(log.GetLevel(), ShouldEqual, log.FatalLevel)
			})

			Convey("should ignore --silent flag if --verbose is set", func() {
				Configure().Run([]string{"pwd", "--silent", "--verbose"})
				So(log.GetLevel(), ShouldEqual, log.DebugLevel)
			})
		})
	})
}
