package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/lyckade/gonetsync/file"
)

//ServerFileGET is a request to the server to send informations
//about a file. If the File does exist a os.FileInfo is sent
//back to the client.
func ServerFileGET(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fp := makeFilePath(vars)
	fi := file.NewFileInfo(fp)
	fi.MakeHash()
	w.Write(fi.JSON())
	myLogger.Println("GET: ", fp)
	r.Body.Close()
	//fmt.Printf("%v", r.FormValue("timestamp"))
	//fmt.Fprintln(w, vars["file"])
	//fmt.Fprintf(w, "Package: %#v; Filepath:%s", vars["package"], //r.FormValue("filepath"))
	//fmt.Fprintln(w, "No")
}

//ServerFilePUT sends a file to the server to store it
func ServerFilePUT(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	/*p := path.Join(
	myConf.BackupFolder,
	vars["package"],
	vars["file"])*/
	fp := makeFilePath(vars)
	// ensure that the directory exists on server
	os.MkdirAll(filepath.Dir(fp), 0777)

	myLogger.Printf("Filepath: %s", fp)
	// Create pointer to new file
	f, err := os.Create(fp)
	if err != nil {
		fmt.Fprintln(w, err)
	}

	// Copy file from request to server
	_, err = io.Copy(f, r.Body)
	f.Close()
	r.Body.Close()
	// Set time
	timets, _ := time.Parse(file.TimestampLayout, r.FormValue("timestamp"))
	err = os.Chtimes(fp, timets, timets)
	//File Info is used for response
	fi := file.NewFileInfo(fp)
	//fi.MakeHash()
	w.Write(fi.JSON())
	if err != nil {
		fmt.Fprintln(w, err)
	}
	

}

func makeFilePath(vars map[string]string) string {
	return path.Join(
		myConf.BackupFolder,
		vars["package"],
		vars["file"])
}
