package main

import "github.com/lyckade/golib/conf"

//ServerConf represents the configuration of the server
type ServerConf struct {
	BackupFolder string `json:"backupFolder"`
	ServerAdress string `json:"serverAdress"`
}

var myConf ServerConf

func init() {
	conf.LoadConf("server.conf.json", myConf, myLogger)
}
