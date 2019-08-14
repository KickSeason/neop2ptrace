package endpoint

import (
	"net"
	"time"

	"github.com/CityOfZion/neo-go/config"

	"neop2ptrace/log"

	"github.com/CityOfZion/neo-go/pkg/network"
	"github.com/CityOfZion/neo-go/pkg/network/payload"
)

type Endpoint struct {
	Addrs     []string
	connected bool
	server    string
	magic     config.NetMode
	conn      net.Conn
	p         network.Peer
	finish    chan<- int
}

var logger = log.NewLogger()

func NewEndpoint(srv string, ch chan<- int) *Endpoint {
	return &Endpoint{
		connected: false,
		server:    srv,
		finish:    ch,
	}
}
func (e *Endpoint) Start() {
	logger.Println("tcp connect", e.server)
	con, err := net.DialTimeout("tcp", e.server, 5*time.Second)
	if err != nil {
		logger.Println(err)
		close(e.finish)
		return
	}
	e.conn = con
	go func(conn net.Conn) {
		for {
			err = e.handleCon(con)
			if err != nil {
				logger.Println(err)
				e.connected = false
				close(e.finish)
				e.Close()
				return
			}
		}
	}(con)
}
func (e *Endpoint) handleMessage(p network.Peer, msg *network.Message) error {
	cmd := msg.CommandType()
	logger.Printf("[endpoint] receive message, type: %s, server: %s", cmd, e.server)
	switch cmd {
	case network.CMDVersion:
		e.magic = msg.Magic
		v := network.NewMessage(msg.Magic, network.CMDVersion, payload.NewVersion(100000, 0, "/Neo:2.10.1/", 0, false))
		vack := network.NewMessage(msg.Magic, network.CMDVerack, nil)
		p.WriteMsg(v)
		p.WriteMsg(vack)
		e.connected = true
		return nil
	case network.CMDAddr:
		addrs := msg.Payload.(*payload.AddressList)
		for _, v := range addrs.Addrs {
			e.Addrs = append(e.Addrs, v.Endpoint.String())
		}
		e.Close()
		// e.SendGetMempool()
		return nil
	case network.CMDVerack:
		if e.connected {
			e.SendGetAddr()
		}
		return nil
	case network.CMDInv:
		hashes := msg.Payload.(*payload.Inventory)
		logger.Println(hashes.Hashes)
		e.Close()
		return nil
	default:
		return nil
	}
}
func (e *Endpoint) handleCon(con net.Conn) error {
	for {
		p := network.NewTCPPeer(con)
		e.p = p
		msg := &network.Message{}
		if err := msg.Decode(con); err != nil {
			return err
		}
		if err := e.handleMessage(p, msg); err != nil {
			return err
		}
	}
}

func (e *Endpoint) SendGetAddr() {
	if !e.connected {
		return
	}
	getaddr := network.NewMessage(e.magic, network.CMDGetAddr, nil)
	e.p.WriteMsg(getaddr)
}

// func (e *Endpoint) SendGetMempool() {
// 	if !e.connected {
// 		return
// 	}
// 	mp := network.NewMessage(e.magic, network.CMDMemPool, nil)
// 	e.p.WriteMsg(mp)
// }

func (e *Endpoint) Close() {
	e.conn.Close()
}
