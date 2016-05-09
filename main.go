package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/server/file", ServerFileGET).Methods("GET")
	router.HandleFunc("/server/file", ServerFilePUT).Methods("PUT")
	router.HandleFunc("/client/file", ClientFileGET).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", router))
}

//ServerFileGET is a request to the server to send informations
//about a file. If the File does exist a os.FileInfo is sent
//back to the client.
func ServerFileGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%#v", r.Body)
	fmt.Fprintln(w, "Here is the file")
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
