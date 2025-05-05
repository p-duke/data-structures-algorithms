package main

import (
	"fmt"
)

/*

Solution
- declare size - len(N) - rows
- hashmap to contain the count of each number in cell
- looping over row major traversal
- if count of cell is > 1 then return false
- reset counts after each row
- looping over col major traversal
- if cell already exists then can return false
- else keep going
- reset after each col

N = The number of rows/columns in the matrix
N = number of rows
M = number of cols
Time Complexity
O((n * m) + k)
Space Complexity
O(n)

*/
func validateSudoku(N [][]int) bool {
	rows := len(N)
	cols := len(N[0])
	count := make(map[int]int)
	for row := 0; row < rows; row++ {
		for k, _ := range count {
			delete(count, k)
		}
		for col := 0; col < cols; col++ {
			cell := N[row][col]
			if cell > rows || cell < 1 {
				return false
			}
			if _, ok := count[cell]; ok {
				return false
			} else {
				count[cell] = 1
			}
		}
	}
	for col := 0; col < cols; col++ {
		for k, _ := range count {
			delete(count, k)
		}
		for row := 0; row < rows; row++ {
			cell := N[row][col]
			if cell > rows || cell < 1 {
				return false
			}
			if _, ok := count[cell]; ok {
				return false
			} else {
				count[cell] = 1
			}
		}
	}
	return true

}
func main() {
	grid1 := [][]int{
		{2, 3, 1},
		{1, 2, 3},
		{3, 1, 2},
	}
	grid2 := [][]int{
		{1, 2, 3},
		{3, 2, 1},
		{3, 1, 2},
	}
	grid3 := [][]int{
		{2, 2, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	grid4 := [][]int{
		{1},
	}
	grid5 := [][]int{
		{-1, -2, -3},
		{-2, -3, -1},
		{-3, -1, -2},
	}
	grid6 := [][]int{
		{1, 3, 3},
		{3, 1, 2},
		{2, 3, 1},
	}
	grid7 := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{1, 3, 2, 4},
		{4, 2, 3, 1},
	}
	grid8 := [][]int{
		{0, 3},
		{3, 0},
	}
	grid9 := [][]int{
		{0, 1},
		{1, 0},
	}

	grid10 := [][]int{
		{1, 1, 6},
		{1, 6, 1},
		{6, 1, 1},
	}
	grid11 := [][]int{
		{1, 2, 3, 4},
		{2, 3, 1, 4},
		{3, 1, 2, 4},
		{4, 2, 3, 1},
	}
	grid12 := [][]int{
		{-1, -2, 12, 1},
		{12, -1, 1, -2},
		{-2, 1, -1, 12},
		{1, 12, -2, -1},
	}
	grid13 := [][]int{
		{2, 3, 3},
		{1, 2, 1},
		{3, 1, 2},
	}
	grid14 := [][]int{
		{1, 3},
		{3, 1},
	}
	grid15 := [][]int{
		{2, 3},
		{3, 2},
	}
	grid16 := [][]int{
		{3, 2},
		{2, 3},
	}
	grid17 := [][]int{
		{2, 3, 1},
		{1, 2, 3},
		{2, 3, 1},
	}
	// Pass
	fmt.Println("Should be true:", validateSudoku(grid1))  // => True
	fmt.Println("Should be false:", validateSudoku(grid2))  // => False
	fmt.Println("Should be false:", validateSudoku(grid3))  // => False
	fmt.Println("Should be true:", validateSudoku(grid4))  // => True
	fmt.Println("Should be false:", validateSudoku(grid5))  // => False
	fmt.Println("Should be false:", validateSudoku(grid6))  // => False
	fmt.Println("Should be false:", validateSudoku(grid7))  // => False
	fmt.Println("Should be false:", validateSudoku(grid8))  // => False
	fmt.Println("Should be false:", validateSudoku(grid9))  // => False
	fmt.Println("Should be false:", validateSudoku(grid10)) // => False
	fmt.Println("Should be false:", validateSudoku(grid11)) // => False
	fmt.Println("Should be false:", validateSudoku(grid12)) // => False
	fmt.Println("Should be false:", validateSudoku(grid13)) // => False
	fmt.Println("Should be false:", validateSudoku(grid14)) // => False
	fmt.Println("Should be false:", validateSudoku(grid15)) // => False
	fmt.Println("Should be false:", validateSudoku(grid16)) // => False
	fmt.Println("Should be false:", validateSudoku(grid17)) // => False
}
