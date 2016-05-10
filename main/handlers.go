package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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

	fmt.Fprintf(w, "%#v", r)

}

//ClientFileGET is the answer of the server to a file request.
//The client decides wether to send the file via put or not.
//Because when it already exists there.
func ClientFileGET(w http.ResponseWriter, r *http.Request) {

}