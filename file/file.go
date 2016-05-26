package file

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lyckade/golib/mylogger"

	"strconv"
	"strings"
	"time"
)

// Info is the datatype for the communiction between client and server
type Info struct {
	//ModTime the last modification time
	ModTime time.Time `json:"modTime"`
	//FilePath is a slice of strings because the file is on the server and
	//the client system
	FilePath []string `json:"filePath"`
	//Exists describes if the file exists
	Exists bool `json:"exists"`
	//Checksum is the md5 hash of the file
	Checksum string `json:"checksum"`
	//fileString is the fpath value when the Object is created.
	fileString string
}

var myLogger = mylogger.NewFileLogger("../log.txt", "")

// TimestampLayout defines the timestamps
const TimestampLayout string = "20060102150405"

// NewFileInfo takes a path as a string and returns a Info struct.
func NewFileInfo(fpath string) Info {
	var fi Info
	fi.fileString = fpath
	Info, err := os.Stat(fpath)
	if err != nil && os.IsNotExist(err) {
		fi.Exists = false
	} else if err != nil {
		myLogger.Print(err)
	} else {
		fi.ModTime = Info.ModTime()
		fi.Exists = true
	}
	fi.FilePath = strings.Split(fpath, strconv.QuoteRune(os.PathSeparator))
	return fi
}

//JSON returns a JSON []byte
func (fi *Info) JSON() []byte {
	s, err := json.Marshal(fi)
	if err != nil {
		myLogger.Print(err)
	}
	return s
}

//Unmarshal takes the JSON Data as []byte and writes it into the
//Info
func (fi *Info) Unmarshal(data []byte) {
	err := json.Unmarshal(data, fi)
	if err != nil {
		myLogger.Print(err)
	}
}

//MakeHash creates the checksum hash
func (fi *Info) MakeHash() {
	h := md5.New()
	f, err := ioutil.ReadFile(fi.fileString)
	if err != nil {
		myLogger.Println(err)
	}
	h.Write(f)
	fi.Checksum = fmt.Sprintf("%x", h.Sum(nil))
	h = md5.New()
}
