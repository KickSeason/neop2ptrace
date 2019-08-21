package nodemap

import (
	"strconv"
	"sync"
)

type Matrix struct {
	mutex  sync.Mutex
	direct bool
	order  []uint64
	m      [][]uint8
}

func NewMatrix(direct bool) Matrix {
	return Matrix{
		direct: direct,
		order:  []uint64{},
		m:      [][]uint8{{}},
	}
}

func (m *Matrix) orderIndex(id uint64) int {
	for i, v := range m.order {
		if v == id {
			return i
		}
	}
	m.order = append(m.order, id)
	return len(m.order) - 1
}
func (m *Matrix) AddEdge(from, to uint64) {
	if !m.direct {
		if from < to {
			from, to = to, from
		}
	}
	m.mutex.Lock()
	fromIndex := m.orderIndex(from)
	toIndex := m.orderIndex(to)

	for len(m.m) <= fromIndex {
		m.m = append(m.m, []uint8{})
	}
	fromRecord := m.m[fromIndex]
	for len(fromRecord) <= toIndex {
		fromRecord = append(fromRecord, uint8(0))
	}
	fromRecord[toIndex] = 1
	m.m[fromIndex] = fromRecord
	m.mutex.Unlock()
}

// func (m Matrix) Edges(from uint64) []uint64 {
// 	result := []uint64{}
// 	fromIndex := m.orderIndex(from)
// 	fromRecord := m.m[fromIndex]
// 	for i, v := range fromRecord {
// 		if v == 1 {
// 			result = append(result, m.order[i])
// 		}
// 	}
// 	return result
// }

func (m Matrix) AllEdges() [][]uint64 {
	result := [][]uint64{}
	for i, v := range m.m {
		from := m.order[i]
		for j, f := range v {
			if f == 1 {
				to := m.order[j]
				pair := []uint64{from, to}
				result = append(result, pair)
			}
		}
	}
	return result
}

func (m Matrix) ToJson() string {
	str := "["
	edges := m.AllEdges()
	for i, v := range edges {
		if 0 < i {
			str += ","
		}
		str += "{"
		str += "\"" + "source" + "\"" + ":"
		str += "\"" + strconv.Itoa(int(v[0])) + "\","
		str += "\"" + "target" + "\"" + ":"
		str += "\"" + strconv.Itoa(int(v[1])) + "\","
		str += "\"" + "value" + "\"" + ":"
		str += "1"
		str += "}"
	}
	str += "]"
	return str
}

func (m Matrix) String() string {
	return m.ToJson()
}
