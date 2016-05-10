package file

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileInfo(t *testing.T) {

	Convey("Should return a JSON string", t, func() {
		fi := NewFileInfo("file_test.go")
		//fmt.Printf("%s", fi.JSON())
		So(string(fi.JSON()), ShouldContainSubstring, `"FilePath":["file_test.go"],"Exists":true}`)
		fi = NewFileInfo("/does/not/exist/file_test.go")
		//fmt.Printf("%s", fi.JSON())
		So(string(fi.JSON()), ShouldContainSubstring, `"Exists":false}`)

	})

}
