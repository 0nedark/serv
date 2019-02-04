package main

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMainPackage(t *testing.T) {
	Convey("package main", t, func() {
		fatalCalled := false
		fatal = func(args ...interface{}) { fatalCalled = true }

		Convey("main function should complete with no errors", func() {
			run = func(args []string) error {
				return nil
			}

			main()
			So(fatalCalled, ShouldBeFalse)
		})

		Convey("main function should log fatal errors", func() {
			run = func(args []string) error {
				return errors.New("test")
			}

			fatal = func(args ...interface{}) {
				So(args[0].(string), ShouldContainSubstring, "test")
			}

			main()
		})
	})
}
