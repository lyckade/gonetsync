package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("fileWalk", t, func() {

		//Create a test server
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//fmt.Printf("TEST: %v", r)
			//myLogger.Printf("Testserver r: %v", r)
			w.WriteHeader(200)
			fmt.Fprintln(w, "Hello, client")
		}))
		defer ts.Close()

		fmt.Println(ts.URL)
		myConf.ServerAdress = ts.URL

		folderWalk(myConf.SyncFolder)

	})

	Convey("CreateUrl inside filewalk", t, func() {
		url := makeURL("http://local:1234", "p1", "a/b", "a/b/c/file.txt")
		So(url.String(), ShouldEqual, "http://local:1234/p1/c/file.txt")

	})

}
