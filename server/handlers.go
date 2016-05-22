package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/gorilla/mux"
)

//ServerFileGET is a request to the server to send informations
//about a file. If the File does exist a os.FileInfo is sent
//back to the client.
func ServerFileGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintln(w, vars["file"])
	fmt.Fprintf(w, "Package: %#v; Filepath:%s", vars["package"], r.FormValue("filepath"))
	//fmt.Fprintln(w, "No")
}

//ServerFilePUT sends a file to the server to store it
func ServerFilePUT(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	p := path.Join(
		myConf.BackupFolder,
		vars["package"],
		vars["file"])
	os.MkdirAll(filepath.Dir(p), 0777)
	fp := p
	myLogger.Printf("Filepath: %s", fp)
	f, err := os.Create(fp)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	defer f.Close()
	h := md5.New()
	_, err = io.Copy(f, r.Body)
	io.Copy(h, r.Body)
	fmt.Fprintf(w, "%v", h.Sum(nil))
	if err != nil {
		fmt.Fprintln(w, err)
	}
	fmt.Fprintf(w, "%#v", fp)
	f.Close()
}
