package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"raspi/apxp/golib/mylogger"

	"github.com/gorilla/mux"
)

//Conf is the configuration structure
type Conf struct {
	BackupFolder string `json:"backupFolder"`
}

var myLogger = mylogger.NewFileLogger("log.txt", "")
var confFile = "../server/server.conf.json"
var conf = loadConf(confFile)

//ServerFileGET is a request to the server to send informations
//about a file. If the File does exist a os.FileInfo is sent
//back to the client.
func ServerFileGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	fmt.Fprintf(w, "Package: %#v; Filepath:%s", vars["package"], r.FormValue("filepath"))

	//fmt.Fprintln(w, "No")
}

//ServerFilePUT sends a file to the server to store it
func ServerFilePUT(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "%#v: %v", r, conf)

}

//ClientFileGET is the answer of the server to a file request.
//The client decides wether to send the file via put or not.
//Because when it already exists there.
func ClientFileGET(w http.ResponseWriter, r *http.Request) {

}

func loadConf(confFile string) Conf {
	var err error
	var b []byte
	var conf Conf
	b, err = ioutil.ReadFile(confFile)
	errorLog(err)
	err = json.Unmarshal(b, &conf)
	errorLog(err)
	return conf
}

func errorLog(err error) {
	if err != nil {
		myLogger.Print(err)
	}
}
