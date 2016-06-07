//Client programm is the client part of the gonetsync
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/lyckade/golib/mylogger"
	"github.com/lyckade/gonetsync/file"
)

var myLog = mylogger.NewFileLogger("client.log", "")

var logLevel = mylogger.DEBUG

var myClient = new(http.Client)

var packageName string

func init() {
	mylogger.Level = mylogger.DEBUG
	flag.StringVar(&packageName, "p", "DefaultPackage", "Package Name")
	//flag.IntVar(&logLevel, "l", myConf.LogLevel, "Define the log level")
	flag.Parse()
	//mylogger.Level = logLevel
}

func main() {
	mylogger.Level = logLevel
	//folderWalk(myConf.SyncFolder, packageName)
	//packageName = myConf.PackageName
	//folderWalk(myConf.SyncFolder, packageName)
	for _, pack := range myPackages {
		folderWalk(pack.SyncFolder, pack.PackageName)
	}
}

func folderWalk(baseFolder, packageName string) {
	filepath.Walk(
		baseFolder,
		func(fpath string, info os.FileInfo, err error) error {
			walkFunc(packageName, baseFolder, fpath, info)
			return err
		})
}

func walkFunc(packageName string, baseFolder string, fpath string, info os.FileInfo) {
	if info.IsDir() == true {
		// Exit when Directory
		return
	}

	// When path is from a file
	ts := makeTimestamp(info)

	//Path transformation
	fpath = filepath.Clean(fpath)

	url := makeURL(myConf.ServerAdress, packageName, baseFolder, fpath, ts)

	// Get fileInfo from Server
	rfiResp, _ := http.Get(url.String())
	body, _ := ioutil.ReadAll(rfiResp.Body)
	rfiResp.Body.Close()
	rfi := new(file.Info)
	rfi.Unmarshal(body)
	//fmt.Println(url.String())
	myLog.Log(mylogger.DEBUG, "GET: ", url.String())
	//myLog.Println("GET: ", url.String())
	if rfi.Exists == true {
		myLog.Log(mylogger.DEBUG, "Exists at server: ", fpath)
		fi := file.NewFileInfo(fpath)
		if rfi.ModTime.After(fi.ModTime) ||
			rfi.ModTime.Equal(fi.ModTime) {
			// Return when local file is younger or equal remote file
			//fmt.Println("Abort file exists")
			return
		}

	}

	//File reading
	fileReader, err1 := os.Open(fpath)
	defer fileReader.Close()
	if err1 != nil {
		myLog.Log(mylogger.ERROR, "Error opening file ", fpath)
		myLog.Log(mylogger.ERROR, err1)
	}

	r, err2 := http.NewRequest("PUT", url.String(), fileReader)
	if err2 != nil {
		myLog.Log(mylogger.ERROR, err2)
	}

	sendClientRequest(r)

}

func makeTimestamp(info os.FileInfo) string {
	ts := info.ModTime().Format(file.TimestampLayout)
	return ts

}

func makeURL(schemeHost, packageStr, baseFolder, filePath, timestamp string) url.URL {
	u, err := url.Parse(schemeHost)
	if err != nil {
		myLog.Log(mylogger.ERROR, err)
	}
	relPath, _ := filepath.Rel(baseFolder, filePath)
	u.Path = packageStr + "/" + filepath.ToSlash(relPath)
	v := url.Values{}
	v.Set("timestamp", timestamp)
	u.RawQuery = v.Encode()
	return *u
}

func makeFileRequest(fpath string, urlStr string) *http.Request {
	//File reading
	fileReader, err1 := os.Open(fpath)
	defer fileReader.Close()
	if err1 != nil {
		myLog.Log(mylogger.ERROR, "Error opening file ", fpath)
		myLog.Log(mylogger.ERROR, err1)
	}

	r, err2 := http.NewRequest("PUT", urlStr, fileReader)
	if err2 != nil {
		myLog.Log(mylogger.ERROR, err2)
	}
	return r
}

func sendClientRequest(r *http.Request) *http.Response {
	resp, err3 := myClient.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	fi := new(file.Info)
	fi.Unmarshal(body)
	defer resp.Body.Close()
	fmt.Printf("RESPONSE: %v\n\n", fi)
	//fmt.Println("RSPERR: ", err3)
	if err3 != nil {
		myLog.Log(mylogger.ERROR, err3)
	}

	return resp
}
