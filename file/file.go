package file

import (
	"encoding/json"
	"os"
	"raspi/apxp/golib/mylogger"
	"strconv"
	"strings"
	"time"
)

// Info is the datatype for the communiction between client and server
type Info struct {
	//ModTime the last modification time
	ModTime time.Time
	//FilePath is a slice of strings because the file is on the server and
	//the client system
	FilePath []string
	//Exists describes if the file exists
	Exists bool
}

var myLogger = mylogger.NewFileLogger("../log.txt", "")

// NewFileInfo takes a path as a string and returns a Info struct.
func NewFileInfo(fpath string) Info {
	var fi Info

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
