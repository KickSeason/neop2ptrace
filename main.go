package main

import (
	"fmt"
	"neop2ptrace/config"
	"neop2ptrace/endpoint"
	"neop2ptrace/log"
	"neop2ptrace/nmap"
	"strings"
)

var seed string
var node *endpoint.Endpoint
var count = 0
var logger = log.NewLogger()
var nodeMap = nmap.NewMap()

func main() {
	config.Load("./config.json")
	seed = config.Seed
	logger.Println("start")
	nodes := []string{seed}
	for i := 0; i < len(nodes); i++ {
		ch := make(chan int)
		node = endpoint.NewEndpoint(nodes[i], ch)
		node.Start()
		if _, ok := <-ch; !ok {
			addrs := filter(node.Addrs)
			fmt.Println(addrs)
			for _, addr := range addrs {
				nodes = append(nodes, addr)
			}
			if 0 < len(addrs) {
				nodeMap.AddNode(nodes[i], addrs)
			}
		}
		fmt.Println(len(nodes))
	}
	fmt.Println(nodeMap.ToJson())
}

func filter(addrs []string) []string {
	rr := []string{}
	for _, addr := range addrs {
		ip := strings.Split(addr, ":")[0]
		if ip != "0.0.0.0" && ip != "127.0.0.1" {
			rr = append(rr, addr)
		}
	}
	return rr
}

func contain(addrs []string, addr string) bool {
	for _, v := range addrs {
		if v == addr {
			return true
		}
	}
	return false
}
