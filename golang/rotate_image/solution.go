package main

// Rotate image in place
func rotate(matrix [][]int) {
	// Get the number of columns and rows
	rows := len(matrix)
	cols := len(matrix[0])

	// Iterate over the original matrix and transpose
	// Use row major traversal
	for r := 0; r < rows; r++ {
		for c := r + 1; c < cols; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}

    // Reverse the rows
    for _, row := range matrix {
        l, r := 0, len(row)-1
        for l < r {
            row[l], row[r] = row[r], row[l]
            l++
            r--
        }
    }
}
