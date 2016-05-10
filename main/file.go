package main

import (
	"encoding/json"
	"os"
	"strconv"
	"strings"
	"time"
)

// FileInfo is the datatype for the communiction between client and server
type FileInfo struct {
	//ModTime the last modification time
	ModTime time.Time
	//FilePath is a slice of strings because the file is on the server and
	//the client system
	FilePath []string
	//Exists describes if the file exists
	Exists bool
}

// NewFileInfo takes a path as a string and returns a FileInfo struct.
func NewFileInfo(fpath string) FileInfo {
	var fi FileInfo

	fileInfo, err := os.Stat(fpath)
	if err != nil && os.IsNotExist(err) {
		fi.Exists = false
	} else if err != nil {
		myLogger.Print(err)
	} else {
		fi.ModTime = fileInfo.ModTime()
		fi.Exists = true
	}
	fi.FilePath = strings.Split(fpath, strconv.QuoteRune(os.PathSeparator))
	return fi
}

//JSON returns a JSON string of the fileInfo
func (fi *FileInfo) JSON() string {
	s, err := json.Marshal(fi)
	if err != nil {
		myLogger.Print(err)
	}

	return string(s)
}
