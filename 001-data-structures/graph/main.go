package main

// Graph represents a graph data structure
type Graph struct {
	// nodes is a map where the keys are strings representing the nodes and the values are slices of strings representing the edges from the node to other nodes
	nodes map[string][]string
}

// NewGraph creates and returns a new Graph instance
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// AddEdge adds an edge from the node "from" to the node "to"
func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
}

func main() {
	// graph stores a new instance of the Graph struct
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("B", "E")
	graph.AddEdge("C", "F")
	graph.AddEdge("C", "G")
}
