package config

import (
	"encoding/json"
	"io/ioutil"

	"neop2ptrace/log"
)

var Servers []string
var logger = log.NewLogger()

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatalln("[config] read config.json failed")
		return
	}
	c := struct {
		Servers []string `json: "servers"`
	}{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		logger.Fatalln("[config] invalid config.json")
		return
	}
	Servers = c.Servers
	logger.Printf("[config] load configuration success. servers: %s", Servers)
}
