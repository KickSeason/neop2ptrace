package main

import (
	"neop2ptrace/config"
	"neop2ptrace/endpoint"

	"neop2ptrace/log"
)

var addrs []string
var servers []*endpoint.Endpoint
var count = 0
var logger = log.NewLogger()

func main() {
	config.Load("./config.json")
	addrs = config.Servers
	servers = make([]*endpoint.Endpoint, len(addrs))
	logger.Println("start")
	start()
	for {
	}
}
func start() {
	for i, v := range addrs {
		logger.Println(v)
		ed := endpoint.NewEndpoint(v)
		servers[i] = ed
		ed.Start()
	}
}
