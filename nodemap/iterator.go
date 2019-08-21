package nodemap

type Iterator struct {
	nmp   *NodeMap
	index int
	end   bool
}

func (nm *NodeMap) Iterator() Iterator {
	it := Iterator{
		nmp:   nm,
		index: 0,
	}
	return it
}

func (it *Iterator) Next() {
	if it.End() {
		return
	}
	it.index++
}

func (it *Iterator) Value() Node {
	n := Node{}
	if !it.End() {
		n = it.nmp.nodes.nodes[it.index]
	}
	return n
}

func (it *Iterator) End() bool {
	if it.nmp.Count() <= it.index {
		it.end = true
	} else {
		it.end = false
	}
	return it.end
}
