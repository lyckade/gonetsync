package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("http routing should work correct", t, func() {
		w := httptest.NewRecorder()
		router := mux.NewRouter()
		addRoutes(router)
		urlValues := url.Values{}
		urlValues.Add("filepath", "eins/zwei/drei")
		req, _ := http.NewRequest("GET",
			"/server/file/myPackage/dsd?"+urlValues.Encode(),
			nil)
		router.ServeHTTP(w, req)
		fmt.Printf("%v", w)
		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Printf("%v", string(body))
	})

}
