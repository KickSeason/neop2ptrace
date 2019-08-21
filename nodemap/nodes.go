package nodemap

import "sync"

type Nodes struct {
	mutex sync.Mutex
	nodes []Node
	nmap  map[uint64]int
}

func NewNodes() Nodes {
	return Nodes{
		nmap: make(map[uint64]int, 1024),
	}
}

func (ns *Nodes) AddNode(n Node) {
	ns.mutex.Lock()
	defer ns.mutex.Unlock()
	if index, ok := ns.nmap[n.ID()]; ok {
		// need to improve
		if ns.nodes[index].group < n.group {
			ns.nodes[index].group = n.group
		}
		return
	}
	ns.nodes = append(ns.nodes, n)
	ns.nmap[n.ID()] = len(ns.nodes) - 1
}

func (ns Nodes) Count() int {
	return len(ns.nodes)
}

func (ns Nodes) ToJson() string {
	str := "["
	for i, v := range ns.nodes {
		if 0 < i {
			str += ","
		}
		str += v.String()
	}
	str += "]"
	return str
}

func (ns Nodes) String() string {
	return ns.ToJson()
}
