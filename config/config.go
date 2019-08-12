package config

import (
	"encoding/json"
	"io/ioutil"

	"neop2ptrace/log"
)

var Seed string
var logger = log.NewLogger()

func Load(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Fatalln("[config] read config.json failed")
		return
	}
	c := struct {
		Seed string `json: "seed"`
	}{}
	err = json.Unmarshal(data, &c)
	if err != nil {
		logger.Fatalln("[config] invalid config.json")
		return
	}
	Seed = c.Seed
	logger.Printf("[config] load configuration success. server: %s", Seed)
}
