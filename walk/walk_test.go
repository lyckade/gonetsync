package walk

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWalk(t *testing.T) {

	Convey("It should walk recursively", t, func() {
		Folder("_test")
	})

}
