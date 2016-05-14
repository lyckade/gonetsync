//Client programm is the client part of the gonetsync
package main

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/lyckade/golib/mylogger"
)

var myLogger = mylogger.NewFileLogger("client.log", "")

const timestampLayout string = "20060102150405"

func main() {
	folderWalk(myConf.SyncFolder)
}

func folderWalk(baseFolder string) {
	filepath.Walk(
		baseFolder,
		func(fpath string, info os.FileInfo, err error) error {
			walkFunc(baseFolder, fpath, info, err)
			return err
		})
}

func walkFunc(baseFolder string, fpath string, info os.FileInfo, err error) error {
	//fmt.Printf("%v\n", fpath)

	//fmt.Printf("%v\n\n", info.ModTime())
	if info.IsDir() != true {

		modTime := info.ModTime()
		ts := modTime.Format(timestampLayout)
		fmt.Println("Timestamp: ", ts)

		urlStr := "http://localhost:8081"
		packageStr := "meinPacket"

		fpath = filepath.Clean(fpath)
		relPath, _ := filepath.Rel(baseFolder, fpath)
		filePath := filepath.ToSlash(relPath)

		fmt.Printf("\nfilePath: %v\n%v\n\n", filePath, fpath)
		fileReader, err1 := os.Open(fpath)
		defer fileReader.Close()
		if err1 != nil {
			myLogger.Println("Error opening file ", fpath)
			myLogger.Println(err)
		}
		urlStr = strings.Join([]string{
			urlStr,
			packageStr,
			filePath}, "/")
		fmt.Printf("URLSTR: %v\n", urlStr)
		r, err2 := http.NewRequest("PUT", urlStr, fileReader)
		if err2 != nil {
			myLogger.Println(err)
		}
		fmt.Println(r)
		client := new(http.Client)
		resp, err3 := client.Do(r)
		fmt.Println("RESPONSE: ", resp)
		fmt.Println("RSPERR: ", err3)
	}
	return err
}
