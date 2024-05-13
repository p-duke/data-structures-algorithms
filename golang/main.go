package main

import (
	"github.com/p-duke/data-structures-algorithms/matrix"
)

func main() {
	m := matrix.Matrix{
		{0, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	result := matrix.SetZeros(m)
	result.Print()
}
