package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lyckade/golib/mylogger"
	"github.com/lyckade/gonetsync/handlers"
)

var myLogger = mylogger.NewFileLogger("log.txt", "")
var confFile = "server.conf.json"

//Conf is the configuration structure
type Conf struct {
	BackupFolder string `json:"backupFolder"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	addRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))

}

func loadConf(confFile string) (Conf, error) {
	var err error
	var b []byte
	var conf Conf
	b, err = ioutil.ReadFile(confFile)
	err = json.Unmarshal(b, &conf)
	return conf, err
}

func addRoutes(router *mux.Router) {
	router.HandleFunc("/server/file/{package}/", handlers.ServerFileGET).Methods("GET")
	router.HandleFunc("/server/file/{package}", handlers.ServerFilePUT).Methods("PUT")
	//	router.HandleFunc("/client/file/{package}", ClientFileGET).Methods("GET")
}
