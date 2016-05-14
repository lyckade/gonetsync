package main

import "github.com/lyckade/golib/conf"

//ClientConf is the structure of the the client json file
type ClientConf struct {
	Port         int    `json:"port"`
	ServerAdress string `json:"serverAdress"`
	SyncFolder   string `json:"syncFolder"`
}

var myConf ClientConf

func init() {
	conf.LoadConf("client.conf.json", &myConf, myLogger)
}
