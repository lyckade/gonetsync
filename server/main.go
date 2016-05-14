//Server handles the requests on the target system.
//The API for the system:
//  URI:/{package}/{filepath}
//The last element of the filepath is always interprated as a filename
//Following methods are supported:
//  * GET: request to get some file informations from the server
//  * PUT: sends a file to the server
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyckade/golib/mylogger"
)

var myLogger = mylogger.NewFileLogger("server.log", "")

func main() {
	router := mux.NewRouter()
	addRoutes(router)
	//fmt.Println(conf.Server.ServerAdress)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func addRoutes(router *mux.Router) {
	router.HandleFunc(`/{package}/{file:.*}`, ServerFileGET).Methods("GET")
	router.HandleFunc(`/{package}/{file:.*}`, ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")
}
