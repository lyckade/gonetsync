package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	Convey("http routing", t, func() {
		w := httptest.NewRecorder()
		router := mux.NewRouter()
		addRoutes(router)

		req, _ := http.NewRequest("GET",
			"/myPackage/dsd/khjsdf/me.txt",
			nil)
		router.ServeHTTP(w, req)
		fmt.Printf("%#v", w)
		//body, _ := ioutil.ReadAll(resp.Body)
		//fmt.Printf("%v", string(body))
	})

}
