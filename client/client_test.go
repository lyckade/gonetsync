package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("fileWalk", t, func() {
		//folderWalk(conf.Client.SyncFolder)
		//folderWalk("C:\\Users\\Schroepfer\\test\\LayoutSemantic")
	})

	Convey("CreateUrl", t, func() {
		url := makeURL("http://local:1234", "p1", "a/b", "a/b/c/file.txt")
		So(url.String(), ShouldEqual, "http://local:1234/p1/c/file.txt")
	})

}
