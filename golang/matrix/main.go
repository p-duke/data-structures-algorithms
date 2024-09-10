package main

import "fmt"

func main() {
	// Example matrix
	matrix := Matrix{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix.Print()

	// DFS
	// rows := len(matrix)
	// cols := len(matrix[0])
	// visited := make([][]bool, rows)
	// for i := range visited {
	// 	visited[i] = make([]bool, cols)
	// }
	// Start DFS from the top-left corner (0, 0)
	// matrix.DFS(0, 0, visited)

	// BFS
	// matrix.BFS()

	// Diagonal Traversal
	// fmt.Println("Primary Diagonal Traversal", matrix.DiagonalTraversal())

	// Spiral Traversal
	fmt.Println("Spiral Traversal", matrix.SpiralTraverse())
}
