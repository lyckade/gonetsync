//gonetsync is a simple programm which allows syncronisation between a server and
//a client inside local networks
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	addRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/server/file/{package}/", ServerFileGET).Methods("GET")
	router.HandleFunc("/server/file/{package}", ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")

}
