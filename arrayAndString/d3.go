package arrayAndString

func setZeroes(matrix [][]int)  {
	rowLen := len(matrix)

	rowMap := map[int]int{}
	colMap := map[int]int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowMap[i] = 1
				colMap[j] = 1
			}
		}
	}

	for row, _ := range rowMap {
		for j :=0; j < len(matrix[row]); j++ {
			matrix[row][j] = 0
		}
	}

	for col, _ := range colMap {
		for i := 0; i < rowLen; i++ {
			matrix[i][col] = 0
		}
	}
}