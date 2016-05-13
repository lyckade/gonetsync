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
	ServerAdress string `json:"serverAdress"`
}

//Server is the conf
var Server ServerConf

//ClientConf is the structure of the the client json file
type ClientConf struct {
	Port         int    `json:"port"`
	ServerAdress string `json:"serverAdress"`
}

//Client is the configuration of the client
var Client ClientConf

func init() {
	LoadConf("../server/server.conf.json", &Server)
	LoadConf("../server/client.conf.json", &Client)
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
