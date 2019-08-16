package config

import (
	"encoding/json"
	"io/ioutil"

	"neop2ptrace/log"
)

var Seed string
var Port int
var logger = log.NewLogger()

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatalln("[config] read config.json failed")
		return
	}
	c := struct {
		Seed string `json: "seed"`
		Port int    `json: "port"`
	}{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		logger.Fatalln("[config] invalid config.json")
		return
	}
	Seed = c.Seed
	Port = c.Port
	logger.Printf("[config] load configuration success. seed: %s, port: %d", Seed, Port)
}
