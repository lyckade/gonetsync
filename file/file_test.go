package file

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileInfo(t *testing.T) {

	Convey("Should return a JSON string", t, func() {
		fi := NewFileInfo("file_test.go")
		//fmt.Printf("%s", fi.JSON())
		So(string(fi.JSON()), ShouldContainSubstring, `"filePath":["file_test.go"],"exists":true`)
		fi.MakeHash()
		fmt.Printf("%v", fi)
	})
	Convey("Should handle if file not exists", t, func() {
		fi := NewFileInfo("/does/not/exist/file_test.go")
		//fmt.Printf("%s", fi.JSON())
		So(string(fi.JSON()), ShouldContainSubstring, `"exists":false`)

	})

}
