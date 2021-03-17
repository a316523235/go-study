package arrayAndString

func rotate(matrix [][]int)  {
	n := len(matrix)
	for i := 0; i < n /2; i++ {
		for j := 0; j < (n+1) / 2; j++ {
			temp := matrix[i][j]
			//matrix[i][j] = 1 左上
			//matrix[j][n - 1 - i] = 1 右上
			//matrix[n-i-1][n-j-1] = 1 右下
			//matrix[n-1-j][i] = 1 左下
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n - 1 - i]
			matrix[j][n - 1 - i] = temp
		}
	}
}