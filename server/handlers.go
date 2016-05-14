package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

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
	vars := mux.Vars(r)
	p := path.Join(
		myConf.BackupFolder,
		vars["package"],
		r.FormValue("filepath"))
	os.MkdirAll(p, 0777)
	fp := path.Join(p, r.FormValue("filename"))
	f, err := os.Create(fp)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	//defer f.Close()
	_, err = io.Copy(f, r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintf(w, "%#v", fp)
	f.Close()
}
