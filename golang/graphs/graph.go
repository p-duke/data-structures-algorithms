package main

import (
	"fmt"
)

type Graph struct {
	adjList map[int][]int
}

// NewGraph initializes an empty graph
func NewGraph() *Graph {
	return &Graph{
		adjList: make(map[int][]int),
	}
}

// AddEdge adds an edge between two vertices (u -> v)
func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	// If it's an undirected graph, add the reverse edge
	// g.adjList[v] = append(g.adjList[v], u)
}

// DFS performs Depth First Search starting from the given vertex
func (g *Graph) DFS(start int) {
	visited := make(map[int]bool)
	fmt.Printf("DFS starting from vertex %d:\n", start)
	g.dfsHelper(start, visited)  // Start DFS
}

// dfsHelper is a recursive function to perform DFS
func (g *Graph) dfsHelper(node int, visited map[int]bool) {
	if visited[node] {
		return // If the node is already visited, return
	}

	// Mark the node as visited and print it
	visited[node] = true
	fmt.Printf("%d ", node)
	fmt.Printf("visited %v\n", visited)

	// Recur for all the adjacent vertices
	for _, neighbor := range g.adjList[node] {
		g.dfsHelper(neighbor, visited)
	}
}

func main() {
	// Create a new graph
	graph := NewGraph()

	// Add edges to the graph
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 5)
	graph.AddEdge(2, 6)
	fmt.Printf("graph %+v\n", graph)
	// Perform DFS starting from vertex 0
	graph.DFS(0)
}


