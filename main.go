package main

import (
	"fmt"
	"neop2ptrace/config"
	"neop2ptrace/endpoint"

	"neop2ptrace/log"
)

var seed string
var node *endpoint.Endpoint
var count = 0
var logger = log.NewLogger()

func main() {
	ch := make(chan int)
	config.Load("./config.json")
	seed = config.Seed
	logger.Println("start")
	node = endpoint.NewEndpoint(seed, ch)
	node.Start()
	if _, ok := <-ch; !ok {
		fmt.Println(node.Addrs)
	}
}
