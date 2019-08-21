package main

import (
	"neop2ptrace/api"
	"neop2ptrace/config"
	"neop2ptrace/endpoint"
	"neop2ptrace/log"
	"neop2ptrace/nodemap"
	"strings"
	"sync"
)

var logger = log.NewLogger()

func main() {
	config.Load("./config.json")
	seed := config.Seed
	logger.Println("start")
	var nodeMap = nodemap.NewNodeMap()
	getAddrs(seed, &nodeMap)
	travel(&nodeMap)
	logger.Println(nodeMap.ToJson())
	srv := api.NewApiServer("127.0.0.1", config.Port, &nodeMap)
	srv.Start()
}

func travel(nm *nodemap.NodeMap) {
	it := nm.Iterator()
	for i := 0; i < nm.Count(); i++ {
		var wg sync.WaitGroup
		for j := 0; !it.End() && j < 10; j++ {
			wg.Add(1)
			addr := it.Value().Address()
			go func() {
				getAddrs(addr, nm)
				wg.Done()
			}()
			it.Next()
			i++
		}
		wg.Wait()
	}

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

func getAddrs(addr string, nm *nodemap.NodeMap) {
	ch := make(chan int)
	ep := endpoint.NewEndpoint(addr, ch)
	ep.Start()
	if _, ok := <-ch; !ok {
		addrs := filter(ep.Addrs)
		node, err := nodemap.NewNode(addr)
		if err != nil {
			logger.Errorln(err)
			return
		}
		peers := []nodemap.Node{}
		for _, v := range addrs {
			n, err := nodemap.NewNode(v)
			if err != nil {
				logger.Errorln(err)
				continue
			}
			peers = append(peers, n)
		}
		nm.AddNode(node, peers)
	}
}
