package repository

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	git "gopkg.in/src-d/go-git.v4"

	log "github.com/sirupsen/logrus"
)

func TestLoadPackage(t *testing.T) {
	Convey("package repository", t, func() {
		Convey("open", func() {
			Convey("should return remote url on success", func() {
				remote, _ := Open("../..")
				So(remote, ShouldEqual, "https://github.com/0nedark/serv.git")
			})

			Convey("should return no error on success", func() {
				_, err := Open("../..")
				So(err, ShouldEqual, nil)
			})

			Convey("should return an error on failure", func() {
				_, err := Open("test_path")
				So(err, ShouldBeError, "repository does not exist")
			})

			Convey("should return empty string on failure", func() {
				remote, _ := Open("test_path")
				So(remote, ShouldEqual, "")
			})
		})

		Convey("clone", func() {
			log.SetLevel(log.PanicLevel)
			Convey("should return no error on success", func() {
				gitClone = func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
					return nil, nil
				}

				msg, _ := Clone("test", "test")
				So(msg, ShouldEqual, "Repository cloned")
			})

			Convey("should return error message on failure", func() {
				gitClone = func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
					return nil, errors.New("test")
				}

				msg, _ := Clone("test", "test")
				So(msg, ShouldEqual, "Failed to clone repository")
			})

			Convey("should return error on failure", func() {
				gitClone = func(path string, isBare bool, o *git.CloneOptions) (*git.Repository, error) {
					return nil, errors.New("test")
				}

				_, err := Clone("test", "test")
				So(err, ShouldNotBeNil)
			})
		})
	})
}
