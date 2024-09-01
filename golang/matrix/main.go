package main

func main() {
	// Example matrix
	matrix := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Create a visited matrix initialized to false
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	// Start DFS from the top-left corner (0, 0)
	matrix.DFS(0, 0, visited)
}
