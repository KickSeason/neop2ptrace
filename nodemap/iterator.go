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
	if it.nmp.Count() <= it.index {
		it.end = true
	}
	return it
}

func (it *Iterator) Next() {
	if it.end {
		return
	}
	it.index++
	if it.nmp.Count() <= it.index {
		it.end = true
		return
	}
}

func (it *Iterator) Value() Node {
	if it.end {
		return Node{}
	}
	if it.nmp.Count() <= it.index {
		it.end = true
		return Node{}
	}
	return it.nmp.nodes.nodes[it.index]
}

func (it *Iterator) End() bool {
	return it.end
}
