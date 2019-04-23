package endpoint

import (
	"net"
	"spamp2p/transaction"
	"time"

	"github.com/CityOfZion/neo-go/config"

	"spamp2p/log"

	"github.com/CityOfZion/neo-go/pkg/network"
	"github.com/CityOfZion/neo-go/pkg/network/payload"
)

type Endpoint struct {
	connected bool
	server    string
	magic     config.NetMode
	conn      net.Conn
	p         network.Peer
}

var logger = log.NewLogger()

func NewEndpoint(srv string) *Endpoint {
	return &Endpoint{
		connected: false,
		server:    srv,
	}
}
func (e *Endpoint) Start() {
	con, err := net.DialTimeout("tcp", e.server, 5*time.Second)
	if err != nil {
		logger.Println(err)
		return
	}
	e.conn = con
	go func(conn net.Conn) {
		for {
			err = e.handleCon(con)
			if err != nil {
				logger.Println(err)
				e.connected = false
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
		v := network.NewMessage(msg.Magic, network.CMDVersion, payload.NewVersion(1234, 0, "/Neo:2.10.1/", 0, false))
		vack := network.NewMessage(msg.Magic, network.CMDVerack, nil)
		p.WriteMsg(v)
		p.WriteMsg(vack)
		e.connected = true
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

func (e *Endpoint) SendTx(tx transaction.TxWrapper) {
	if !e.connected {
		return
	}
	//logger.Printf("[Endpoint] send tx. srv: %s\n", e.server)
	txmsg := network.NewMessage(e.magic, network.CMDTX, tx)
	e.p.WriteMsg(txmsg)
}
