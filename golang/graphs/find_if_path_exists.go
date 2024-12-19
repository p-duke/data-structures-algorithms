type Graph struct {
	adjList map[int][]int
}

func (g *Graph) AddEdge(u, v int) {
	g.adjList[u] = append(g.adjList[u], v)
	g.adjList[v] = append(g.adjList[v], u)
}

func (g *Graph) DFS(visited map[int]bool, vertex, dest int) bool {
    if visited[vertex] {
        return false
    }

    visited[vertex] = true

    if vertex == dest {
        return true
    }

    for _, neighbor := range g.adjList[vertex] {
        if g.DFS(visited, neighbor, dest) {
            return true
        }
    }

    return false
}

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := &Graph{adjList: make(map[int][]int)}
	for _, v := range edges {
		graph.AddEdge(v[0], v[1])
	}

    visited := make(map[int]bool)

    return graph.DFS(visited, source, destination)
}
