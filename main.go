package main

import (
	"os"
	"spamp2p/config"
	"spamp2p/endpoint"
	"spamp2p/transaction"
	"strconv"
	"time"

	"spamp2p/log"
)

var addrs []string
var servers []*endpoint.Endpoint
var count = 0
var logger = log.NewLogger("main")

func main() {
	if len(os.Args) < 2 {
		logger.Fatalln("please input tps.")
		return
	}
	tps, err := strconv.Atoi(os.Args[1])
	if err != nil {
		logger.Fatalln("tps invalid.")
		return
	}
	var con, interval int
	if 1000 < tps {
		con = tps / 1000
		interval = 1
	} else {
		interval = 1000 / tps
		con = 1
	}
	logger.Printf("tps: %d, concorrent: %d, interval: %dms.\n", tps, con, interval)
	config.Load("./config.json")
	addrs = config.Servers
	servers = make([]*endpoint.Endpoint, len(addrs))
	start()
	go statastic()
	timer := time.NewTicker(time.Duration(interval) * time.Millisecond)
	for {
		select {
		case <-timer.C:
			sendtx(con)
		}
	}
}

func sendtx(con int) {
	for i := 0; i < con; i++ {
		tx := transaction.InvokeTx()
		count++
		for _, ed := range servers {
			go ed.SendTx(tx)
		}
	}
}

func start() {
	for i, v := range addrs {
		ed := endpoint.NewEndpoint(v)
		servers[i] = ed
		ed.Start()
	}
}

func statastic() {
	interval := 15
	last := 0
	start := time.Now()
	timer := time.NewTicker(time.Duration(interval) * time.Second)
	for {
		select {
		case <-timer.C:
			increase := count - last
			last = count
			logger.Printf("tps now %0.2f\n", float64(increase)/float64(interval))
			now := time.Now()
			period := now.Sub(start).Seconds()
			logger.Printf("tps summary %0.2f\n", float64(count)/period)
		}
	}
}
