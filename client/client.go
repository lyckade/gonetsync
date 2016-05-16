//Client programm is the client part of the gonetsync
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"

	"github.com/lyckade/golib/mylogger"
)

var myLogger = mylogger.NewFileLogger("client.log", "")

type doer interface {
	Do(*http.Request) (*http.Response, error)
}

var myClient = new(http.Client)

const timestampLayout string = "20060102150405"

func main() {

	folderWalk(myConf.SyncFolder)
}

func folderWalk(baseFolder string) {
	filepath.Walk(
		baseFolder,
		func(fpath string, info os.FileInfo, err error) error {
			walkFunc(baseFolder, fpath, info)
			return err
		})
}

func walkFunc(baseFolder string, fpath string, info os.FileInfo) {

	if info.IsDir() != true {
		ts, _ := makeTimestamp(info)
		fmt.Println("Timestamp: ", ts)

		packageStr := "meinPacket"

		//Path transformation
		fpath = filepath.Clean(fpath)

		url := makeURL(myConf.ServerAdress, packageStr, baseFolder, fpath)
		fmt.Println(url.String())
		r := makeFileRequest(fpath, url.String())
		//fmt.Printf("%v", r)

		sendClientRequest(r)
	}

}

func makeTimestamp(info os.FileInfo) (int, error) {
	ts := info.ModTime().Format(timestampLayout)
	return strconv.Atoi(ts)
}

func makeURL(schemeHost string, packageStr string, baseFolder string, filePath string) url.URL {
	url, err := url.Parse(schemeHost)
	if err != nil {
		myLogger.Print(err)
	}
	relPath, _ := filepath.Rel(baseFolder, filePath)
	url.Path = packageStr + "/" + filepath.ToSlash(relPath)
	return *url
}

func makeFileRequest(fpath string, urlStr string) *http.Request {
	//File reading
	fileReader, err1 := os.Open(fpath)
	defer fileReader.Close()
	if err1 != nil {
		myLogger.Println("Error opening file ", fpath)
		myLogger.Println(err1)
	}

	r, err2 := http.NewRequest("PUT", urlStr, fileReader)
	if err2 != nil {
		myLogger.Println(err2)
	}
	return r
}

func sendClientRequest(r *http.Request) {
	resp, err3 := myClient.Do(r)
	fmt.Println("RESPONSE: ", resp)
	fmt.Println("RSPERR: ", err3)
}
