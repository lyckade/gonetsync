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

var myLogger = mylogger.NewFileLogger("client.log", "")

var myClient = new(http.Client)

var packageName string

func init() {
	flag.StringVar(&packageName, "p", "DefaultPackage", "Package Name")
	flag.Parse()
}

func main() {

	//folderWalk(myConf.SyncFolder, packageName)
	folderWalk(myConf.SyncFolder, packageName)
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
	fmt.Println("Timestamp: ", ts)

	//Path transformation
	fpath = filepath.Clean(fpath)

	url := makeURL(myConf.ServerAdress, packageName, baseFolder, fpath, ts)
	fmt.Println(url.String())

	// Get fileInfo from Server
	rfiResp, _ := http.Get(url.String())
	body, _ := ioutil.ReadAll(rfiResp.Body)
	rfi := new(file.Info)
	rfi.Unmarshal(body)
	myLogger.Println("GET: ", url.String())
	if rfi.Exists == true {
		myLogger.Println("Exists at server: ", fpath)
		return
		//fi := file.NewFileInfo(fpath)
		/*fi.MakeHash()
		if fi.Checksum == rfi.Checksum {
			myLogger.Println("Same file at server: ", fpath)
			return
		}*/

	}

	//File reading
	fileReader, err1 := os.Open(fpath)
	defer fileReader.Close()
	if err1 != nil {
		myLogger.Println("Error opening file ", fpath)
		myLogger.Println(err1)
	}

	r, err2 := http.NewRequest("PUT", url.String(), fileReader)
	if err2 != nil {
		myLogger.Println(err2)
	}

	sendClientRequest(r)

}

func makeTimestamp(info os.FileInfo) string {
	return info.ModTime().Format(file.TimestampLayout)
	
}

func makeURL(schemeHost, packageStr, baseFolder, filePath, timestamp string) url.URL {
	u, err := url.Parse(schemeHost)
	if err != nil {
		myLogger.Print(err)
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
		myLogger.Println("Error opening file ", fpath)
		myLogger.Println(err1)
	}

	r, err2 := http.NewRequest("PUT", urlStr, fileReader)
	if err2 != nil {
		myLogger.Println(err2)
	}
	return r
}

func sendClientRequest(r *http.Request) *http.Response {
	resp, err3 := myClient.Do(r)
	body, _ := ioutil.ReadAll(resp.Body)
	fi := new(file.Info)
	fi.Unmarshal(body)
	defer resp.Body.Close()
	fmt.Printf("RESPONSE: %v", fi)
	//fmt.Println("RSPERR: ", err3)
	if err3 != nil {
		myLogger.Println(err3)
	}

	return resp
}
