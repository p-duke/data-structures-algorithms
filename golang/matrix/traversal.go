package matrix

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

func (m Matrix) Print() {
	for _, v := range m {
		fmt.Println(v)
	}
}
