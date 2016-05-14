package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyckade/golib/mylogger"
)

var myLogger = mylogger.NewFileLogger("server.log", "")

func main() {
	router := mux.NewRouter().StrictSlash(true)
	addRoutes(router)
	//fmt.Println(conf.Server.ServerAdress)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/server/file/{package}/", ServerFileGET).Methods("GET")
	router.HandleFunc("/server/file/{package}", ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")
}
