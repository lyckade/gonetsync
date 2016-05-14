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
	router.HandleFunc("/{package}/{file:.*}", ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")
}
