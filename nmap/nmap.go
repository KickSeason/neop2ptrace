package nmap

type NodeMap struct {
	nmap map[string][]string
}

func NewMap() NodeMap {
	return NodeMap{
		nmap: make(map[string][]string, 1024),
	}
}

func (nm *NodeMap) AddNode(node string, peers []string) bool {
	nm.nmap[node] = peers
	return true
}

func (nm *NodeMap) RemoveNode(node string) bool {
	return true
}

func (nm *NodeMap) ToJson() string {
	str := "{"
	j := 0
	for k, v := range nm.nmap {
		if 0 < j {
			str += ","
		}
		str += "\"" + k + "\""
		str += ": [ "
		for i, e := range v {
			if 0 < i {
				str += ", "
			}
			str += "\"" + e + "\""
		}
		str += "]"
		j++
	}
	return str
}

func (nm *NodeMap) String() string {
	return nm.ToJson()
}

// func arrayExcept(ar, br []string) ([]string, []string) {
// 	kr := make([]string, 0, len(ar))
// 	dr := make([]string, 0, len(ar))
// 	for _, va := range br {
// 		keep := true
// 		for _, vb := range ar {
// 			if va == vb {
// 				keep = false
// 				break
// 			}
// 		}
// 		if keep {
// 			kr = append(kr, va)
// 		} else {
// 			dr = append(dr, va)
// 		}
// 	}
// }
