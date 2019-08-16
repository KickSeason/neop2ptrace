package nodemap

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Node struct {
	id    uint64
	ip    net.IP
	port  uint16
	group uint16
}

func id(ip net.IP, port uint16) uint64 {
	parr := make([]byte, 2)
	binary.LittleEndian.PutUint16(parr, port)
	idarr := append(ip.To4(), parr...)
	idarr = append(idarr, []byte{0, 0}...)
	id := binary.LittleEndian.Uint64(idarr)
	return id
}

func NewNode(addr string) (Node, error) {
	iport := strings.Split(addr, ":")
	port, err := strconv.Atoi(iport[1])
	if err != nil {
		return Node{}, errors.New(fmt.Sprintf("Failed new node, string: %s. Wrong port", addr))
	}
	ip := net.ParseIP(iport[0])
	if ip == nil {
		return Node{}, errors.New(fmt.Sprintf("Failed new node, string: %s. Wrong ip", addr))
	}
	return Node{
		id:   id(ip, uint16(port)),
		ip:   net.ParseIP(iport[0]),
		port: uint16(port),
	}, nil
}

func (n Node) ID() uint64 {
	return n.id
}

func (n Node) Address() string {
	return n.ip.String() + ":" + strconv.Itoa(int(n.port))
}

func (n Node) ToJson() string {
	str := "{"
	str += "\"" + "id" + "\"" + ":"
	str += "\"" + strconv.Itoa(int(n.id)) + "\","
	str += "\"" + "group" + "\"" + ":"
	str += strconv.Itoa(int(n.group))
	str += "}"
	return str
}

func (n Node) String() string {
	return n.ToJson()
}
