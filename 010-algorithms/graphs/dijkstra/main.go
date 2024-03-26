package main

import "fmt"

// Graph represents a graph with nodes and edges
type Graph struct {
	nodes map[string][]string
}

// NewGraph creates and returns a new Graph instance
func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]string)}
}

// AddEdge adds an edge between two nodes in the graph
func (g *Graph) AddEdge(from, to string) {
	g.nodes[from] = append(g.nodes[from], to)
}

// Dijkstra implements Dijkstra's algorithm to find the shortest path from the startNode to all other nodes in the graph
func (g *Graph) Dijkstra(startNode string) map[string]int {
	// distances is a map that stores the shortest distance from the startNode to each node in the graph
	distances := make(map[string]int)

	// predecessors is a map that stores the predecessor of each node in the shortest path from the startNode
	predecessors := make(map[string]string)

	// Initialize the distances and predecessors maps
	for node := range g.nodes {
		distances[node] = 1_000_000
		if node == startNode {
			distances[node] = 0
		}
		predecessors[node] = ""
	}

	// activeNodes is a slice that contains all nodes that have not yet been processed by the algorithm
	activeNodes := make([]string, 0, len(g.nodes))
	for node := range g.nodes {
		activeNodes = append(activeNodes, node)
	}

	// Loop through the activeNodes slice until it is empty
	for len(activeNodes) > 0 {
		// Find the node with the smallest distance in the activeNodes slice
		currentMin := 1_000_000
		var currentNode string
		for _, node := range activeNodes {
			if distances[node] < currentMin {
				currentMin = distances[node]
				currentNode = node
			}
		}

		// Remove the current node from the activeNodes slice
		activeNodes = removeElement(activeNodes, currentNode)

		// Update the distances and predecessors maps for the neighbors of the current node
		for _, neighbor := range g.nodes[currentNode] {
			alt := distances[currentNode] + 1
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				predecessors[neighbor] = currentNode
				activeNodes = append(activeNodes, neighbor)
			}
		}
	}

	// Return the distances map
	return distances
}

// removeElement removes an element from a slice
func removeElement(slice []string, element string) []string {
	for i, s := range slice {
		if s == element {
			slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	return slice
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge("A", "B")
	graph.AddEdge("A", "C")
	graph.AddEdge("B", "D")
	graph.AddEdge("B", "E")
	graph.AddEdge("C", "F")
	graph.AddEdge("C", "G")

	// Run Dijkstra's algorithm on the graph starting from node "A".
	shortestPaths := graph.Dijkstra("A")

	// Print the shortest path distances from node "A" to all other nodes.
	for node, distance := range shortestPaths {
		fmt.Printf("%s: %d\n", node, distance)
	}
}
