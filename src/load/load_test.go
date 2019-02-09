package load

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func readFileMock(filecontent string, err error) func(file string) ([]byte, error) {
	return func(file string) ([]byte, error) {
		return []byte(filecontent), err
	}
}

func TestLoadPackage(t *testing.T) {
	Convey("package load", t, func() {
		Convey("on success", func() {
			Convey("should return object with order", func() {
				read = readFileMock("order: [test]", nil)
				obj, _ := Config("test.yml")
				So(obj.Order, ShouldResemble, []string{"test"})
			})

			Convey("should return object with groups", func() {
				read = readFileMock("{\"groups\":{}}", nil)
				obj, _ := Config("test.yml")
				So(obj.Groups, ShouldResemble, map[string][]Service{})
			})
		})

		Convey("on error", func() {
			Convey("should return file read error", func() {
				read = readFileMock("", errors.New("error"))
				_, err := Config("test.yml")
				So(err.Error(), ShouldContainSubstring, "error")
			})

			Convey("should return yaml parse error", func() {
				read = readFileMock("asd", nil)
				_, err := Config("test.yml")
				So(err.Error(), ShouldContainSubstring, "yaml")
			})
		})
	})
}
