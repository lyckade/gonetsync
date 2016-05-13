package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/lyckade/golib/mylogger"
	"github.com/lyckade/gonetsync/conf"
)

var myLogger = mylogger.NewFileLogger("../log.txt", "")

func main() {
	folderWalk(conf.Client.SyncFolder)
}

func folderWalk(f string) {
	filepath.Walk(f, func(fpath string, info os.FileInfo, err error) error {
		fmt.Printf("%v\n", fpath)
		//fmt.Printf("%v\n", info.IsDir())
		//fmt.Printf("%v\n\n", info.ModTime())
		if info.IsDir() != true {

			/*urlStr := path.Join(
			"localhost:8081",
			"server",
			"file",
			"dasPacket")*/

			fpath = filepath.Clean(fpath)
			urlStr := "http://localhost:8081/server/file/testpa"
			urlValues := url.Values{}

			urlValues.Add("filepath", filepath.Dir(fpath))
			urlValues.Add("filename", info.Name())
			urlStr = urlStr + "?" + urlValues.Encode()
			fileReader, err1 := os.Open(fpath)
			defer fileReader.Close()
			if err1 != nil {
				myLogger.Println(err)
			}
			fmt.Printf("URLSTR: %v\n", urlStr)
			r, err2 := http.NewRequest("PUT", urlStr, fileReader)
			if err2 != nil {
				myLogger.Println(err)
			}
			//fmt.Println(r)
			client := new(http.Client)
			resp, err3 := client.Do(r)
			fmt.Println("RESPONE: ", resp)
			fmt.Println("RSPERR: ", err3)
		}

		return err
	})
}
