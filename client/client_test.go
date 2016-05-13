package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("fileWalk", t, func() {
		//folderWalk(conf.Client.SyncFolder)
		folderWalk("../walk/_test")
	})

}
