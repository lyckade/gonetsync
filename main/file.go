package main

import "os"

// FileInfo is the datatype for the communiction between client and server
type FileInfo struct {
	//FileInfo from the os package for more information about the file
	os.FileInfo
	//FilePath is a slice of strings because the file is on the server and
	//the client system
	FilePath []string
	//Exists describes if the file exists
	Exists bool
}

// NewFileInfo takes a path as a string and returns a FileInfo struct.
func NewFileInfo(fpath string) FileInfo {
	var fi FileInfo
	return fi
}
