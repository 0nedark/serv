package verify

import (
	"errors"
	"testing"

	"github.com/0nedark/serv/src/load"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVerifyPackage(t *testing.T) {
	Convey("package verify", t, func() {
		openCalled := false
		open = func(path string) (string, error) {
			openCalled = true
			return "", nil
		}

		cloneCalled := false
		clone = func(url, path string) { cloneCalled = true }

		Convey("Each", func() {
			groups := load.Groups{}

			Convey("should not try to open service with no repository", func() {
				groups["test"] = load.Services{load.Service{}}
				Each(groups)
				So(openCalled, ShouldBeFalse)
			})

			Convey("should try to open service with repository", func() {
				repository := load.Repository{URL: "git repo url", Path: ".."}
				groups["test"] = load.Services{load.Service{Repository: repository}}
				Each(groups)
				So(openCalled, ShouldBeTrue)
			})

			Convey("should try to clone service if it doesn't exist locally", func() {
				repository := load.Repository{URL: "git repo url", Path: ".."}
				groups["test"] = load.Services{load.Service{Repository: repository}}
				open = func(path string) (string, error) { return "", errors.New("") }
				Each(groups)
				So(cloneCalled, ShouldBeTrue)
			})

			Convey("should not try to clone service if no repository was provided", func() {
				groups["test"] = load.Services{load.Service{}}
				open = func(path string) (string, error) { return "", errors.New("") }
				Each(groups)
				So(cloneCalled, ShouldBeFalse)
			})
		})
	})
}
