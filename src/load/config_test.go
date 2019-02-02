package load

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func readFileMock(filecontent string, err error) fileReader {
	return func(file string) ([]byte, error) {
		return []byte(filecontent), err
	}
}

func TestConfigLoading(t *testing.T) {
	Convey("Configuration is loading correctly", t, func() {
		Convey("order is loaded", func() {
			obj, _ := config("test.yml", readFileMock("order: [test]", nil))
			So(obj.Order, ShouldResemble, []string{"test"})
		})

		Convey("groups are loaded", func() {
			obj, _ := config("test.yml", readFileMock("{\"groups\":{}}", nil))
			So(obj.Groups, ShouldResemble, map[string][]Service{})
		})

		Convey("return file read error", func() {
			_, err := config("test.yml", readFileMock("", errors.New("error")))
			So(err.Error(), ShouldContainSubstring, "error")
		})

		Convey("return yaml parse error", func() {
			_, err := config("test.yml", readFileMock("asd", nil))
			So(err.Error(), ShouldContainSubstring, "yaml")
		})
	})
}
