package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyckade/golib/mylogger"
	"github.com/lyckade/gonetsync/handlers"
)

var myLogger = mylogger.NewFileLogger("log.txt", "")

func main() {
	router := mux.NewRouter().StrictSlash(true)
	addRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/server/file/{package}/", handlers.ServerFileGET).Methods("GET")
	router.HandleFunc("/server/file/{package}", handlers.ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")
}
