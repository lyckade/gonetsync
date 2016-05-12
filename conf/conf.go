// Package conf represents the conf of the whole gonetsync app
package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/lyckade/golib/mylogger"
)

//Logger is the global logger
var Logger = mylogger.NewFileLogger("../log.txt", "")

//ServerConf represents the configuration of the server
type ServerConf struct {
	BackupFolder string `json:"backupFolder"`
}

//Server is the conf
var Server ServerConf

//NewServer creates a new server configuration object
func NewServer() ServerConf {
	var server ServerConf
	LoadConf("../server/server.conf.json", &server)
	return server
}

func init() {
	LoadConf("../server/server.conf.json", &Server)
}

//LoadConf loads the configuration from a json file into a
//struct
func LoadConf(confFile string, confStruc interface{}) {
	var err error
	var b []byte

	b, err = ioutil.ReadFile(confFile)
	errorLog(err)
	err = json.Unmarshal(b, &confStruc)
	errorLog(err)

}

func errorLog(err error) {
	if err != nil {
		Logger.Print(err)
	}
}
