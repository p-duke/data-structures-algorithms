package matrix

func SetZeros(matrix Matrix) Matrix {
	colLength := len(matrix[0])
	rowLength := len(matrix)

	markCol := make([]int, colLength)
	markRow := make([]int, rowLength)

	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if matrix[row][col] == 0 {
				markCol[col] = 1
				markRow[row] = 1
			}
		}

	}

	for row := 0; row < rowLength; row++ {
		for col := 0; col < colLength; col++ {
			if markRow[row] == 1 || markCol[col] == 1 {
				matrix[row][col] = 0
			}
		}
	}

	return matrix
}
