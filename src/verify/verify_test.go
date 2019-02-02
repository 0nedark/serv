package verify

import (
	"testing"

	"github.com/0nedark/serv/src/load"
	. "github.com/smartystreets/goconvey/convey"
)

func TestVerifyPackage(t *testing.T) {
	Convey("package verify", t, func() {
		Convey("emptyRepository", func() {
			Convey("should return true if repository structure is empty", func() {
				So(emptyRepository(load.Repository{}), ShouldBeTrue)
			})

			Convey("should return false if repository structure is not empty", func() {
				repository := load.Repository{URL: "git repo url", Path: ".."}
				So(emptyRepository(repository), ShouldBeFalse)
			})
		})

		Convey("selectRepositories", func() {
			repositories := make([]load.Repository, 0)

			Convey("should select only non empty repositories", func() {
				services := load.Services{
					load.Service{Repository: load.Repository{}},
					load.Service{Repository: load.Repository{URL: "git repo url", Path: ".."}},
				}

				filtered := selectRepositories(services, repositories)
				So(filtered[0], ShouldResemble, services[1].Repository)
			})

			Convey("should remove everything if no repositories are found", func() {
				services := load.Services{load.Service{}, load.Service{}}
				filtered := selectRepositories(services, repositories)
				So(filtered, ShouldBeEmpty)
			})

			Convey("should return empty array on empty service array", func() {
				services := load.Services{}
				filtered := selectRepositories(services, repositories)
				So(filtered, ShouldBeEmpty)
			})
		})
	})
}
