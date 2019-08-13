package map


type NodeMap struct {
	nmap  map[string][]string
}

func NewMap() NodeMap {
	return NodeMap{
		nmap: make(map[string][]string, 1024),
	}
}

func (nm *NodeMap)AddNode(node string, peers []string) bool {
	if n, ok := nm[node]; !ok {
		nm[node] = peers
		return true
	}
	for _, peer := range nm[node] {

	}
}

func (nm *NodeMap)RemoveNode(node string) bool {
	return true
}

func ArrayExcept(ar, br []string) []string {
	tr := make([]string, len(ar))
	for _, va := range br {
		for _, vb := range ar {
			
		}
	}
}