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

// There are two separate ways to traverse a matrix diagonally
// Primary - Top-left to bottom right
// Secondary - Top-right to bottom-left
//
// You can also traverse an array using either zigzag traversal or straight diagonal traversal
// Below is a straight diagonal traversal example
func (m Matrix) DiagonalTraversal() []int {
	if len(m) == 0 {
		return []int{}
	}

	result := make([]int, 0)
	rows := len(m)
	cols := len(m[0])

	// Traverse starting from first row
	for col := 0; col < cols; col++ {
		i, j := 0, col
		// Traverse each diagonal
		for i < rows && j >= 0 {
			result = append(result, m[i][j])
			i++
			j--
		}
	}

	// Traverse from the first column skipping the first element to avoid duplication
	for row := 1; row < rows; row++ {
		i, j := row, cols-1
		for i < rows && j >= 0 {
			result = append(result, m[i][j])
			i++
			j--
		}
	}

	return result
}

func (m Matrix) SpiralTraverse() []int {
	result := []int{}
	if len(m) == 0 {
		return result
	}

	top, bottom := 0, len(m)-1
	left, right := 0, len(m[0])-1

	// Traverse in spiral order
	for top <= bottom && left <= right {
		// Traverse from left to right across the top row
		for i := left; i <= right; i++ {
			result = append(result, m[top][i])
		}

		top++

		// Traverse from top to bottom down the right column
		for i := top; i <= bottom; i++ {
			result = append(result, m[i][right])
		}

		right--

		// Traverse from the right to left across the bottom row
		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, m[bottom][i])
			}

			bottom--
		}

		// Traverse from bottom to top up the left column
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result, m[i][left])
			}

			left++
		}
	}

	return result
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

func (m Matrix) BFS() {
	if len(m) == 0 {
		return
	}

	rows := len(m)
	cols := len(m)

	// Directions array to help explore neighbors (up, down, left, right)
	directions := [][]int{
		{-1, 0}, // up
		{1, 0},  // down
		{0, -1}, // left
		{0, 1},  // right
	}

	// Queue to manage BFS
	queue := [][]int{}

	// Starting BFS from top-left corner
	queue = append(queue, []int{0, 0})

	// Visited matrix to keep track of visited cells
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	visited[0][0] = true

	for len(queue) > 0 {
		// Dequeue a cell from the front
		cell := queue[0]
		queue = queue[1:]

		row, col := cell[0], cell[1]
		fmt.Printf("Visiting cell: [%d, %d] with value : %d\n", row, col, m[row][col])

		// Explore all 4 possible directions
		for _, direction := range directions {
			newRow, newCol := row+direction[0], col+direction[1]

			// Check if the new cell is within bounds and not yet visited
			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && !visited[newRow][newCol] {
				visited[newRow][newCol] = true
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
}

func (m Matrix) Transpose() [][]int {
	// Get the number of columns and rows
	rows := len(m)
	cols := len(m[0])

	// Create a new matrix with dimensions cols x rows
	result := make([][]int, cols)
	for i := range result {
		result[i] = make([]int, rows)
	}

	// Iterate over the original matrix and assign values to the new transposed matrix
	// Use row major traversal
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			result[c][r] = m[r][c]
		}
	}

	return result

}

func (m Matrix) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}
