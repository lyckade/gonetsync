package main

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileInfo(t *testing.T) {

	Convey("Should return a JSON string", t, func() {
		fi := NewFileInfo("file_test.go")
		fmt.Printf("%s", fi.JSON())
		fi = NewFileInfo("/does/not/exist/file_test.go")
		fmt.Printf("%s", fi.JSON())
	})

}
