package load

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func readFileMock(filecontent string, err error) ReadFileFunc {
	return func(file string) ([]byte, error) {
		return []byte(filecontent), err
	}
}

func TestLoadPackage(t *testing.T) {
	Convey("package load", t, func() {
		Convey("on success", func() {
			Convey("should return object with order", func() {
				obj, _ := NewConfig("test.yml", readFileMock("order: [test]", nil))
				So(obj.Order, ShouldResemble, []string{"test"})
			})

			Convey("should return object with groups", func() {
				obj, _ := NewConfig("test.yml", readFileMock("{\"groups\":{}}", nil))
				So(obj.Groups, ShouldResemble, map[string][]Service{})
			})
		})

		Convey("on error", func() {
			Convey("should return file read error", func() {
				_, err := NewConfig("test.yml", readFileMock("", errors.New("error")))
				So(err.Error(), ShouldContainSubstring, "error")
			})

			Convey("should return yaml parse error", func() {
				_, err := NewConfig("test.yml", readFileMock("asd", nil))
				So(err.Error(), ShouldContainSubstring, "yaml")
			})
		})
	})
}
