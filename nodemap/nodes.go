package nodemap

type Nodes struct {
	nodes []Node
	nmap  map[uint64]int
}

func NewNodes() Nodes {
	return Nodes{
		nmap: make(map[uint64]int, 1024),
	}
}

func (ns *Nodes) AddNode(n Node) {
	if index, ok := ns.nmap[n.ID()]; ok {
		ns.nodes[index] = n
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
