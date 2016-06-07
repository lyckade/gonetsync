package main

import "github.com/lyckade/golib/conf"

//ClientConf is the structure of the the client json file
type ClientConf struct {
	Port         int    `json:"port"`
	LogLevel     int    `json:"logLevel"`
	ServerAdress string `json:"serverAdress"`
	SyncFolder   string `json:"syncFolder"`
	PackageName  string `json:"package"`
}

//ClientPackages is the structure of the packages.json file
type ClientPackages []struct {
	SyncFolder  string `json:"syncFolder"`
	PackageName string `json:"package"`
}

var myConf ClientConf
var myPackages ClientPackages

func init() {
	conf.LoadConf("client.conf.json", &myConf, myLog)
	conf.LoadConf("packages.json", &myPackages, myLog)
}
