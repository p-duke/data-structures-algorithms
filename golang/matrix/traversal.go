package main

import "fmt"

type Matrix [][]int

func (m Matrix) ColumnMajorTraversal() {
	for col := 0; col < len(m[0]); col++ {
		fmt.Printf("Col %d\n", col)

		for row := 0; row < len(m); row++ {
			fmt.Println(m[row][col])
		}
	}
}

func (m Matrix) RowMajorTraversal() {
	for row := 0; row < len(m); row++ {
		fmt.Printf("Row %d\n", row)

		for col := 0; col < len(m[row]); col++ {
			fmt.Println(m[row][col])
		}
	}
}

// Up - (row - 1)
// Down - (row + 1)
// Left - (col - 1)
// Right - (col + 1)
func (m Matrix) DFS(r, c int, visited [][]bool) {
	rows := len(m)
	cols := len(m[0])

	// Base case
	if r < 0 || c < 0 || r >= rows || c >= cols || visited[r][c] || m[r][c] == 0 {
		return
	}

	visited[r][c] = true
	fmt.Printf("Visited cell r: %d c: %d val: %d\n", r, c, m[r][c])

	m.DFS(r-1, c, visited) // Up
	m.DFS(r+1, c, visited) // Down
	m.DFS(r, c+1, visited) // Left
	m.DFS(r, c-1, visited) // Right
}

func (m Matrix) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}
