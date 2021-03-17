package arrayAndString

func findDiagonalOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return res
	}

	m, n:= len(matrix), len(matrix[0])
	listMap := map[int][]int{}

	for line := 0; line <= (m- 1) + (n - 1) ; line++  {
		listMap[line] = []int{}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			listMap[i+j] = append(listMap[i+j], matrix[i][j])
		}
	}

	for line := 0; line <= (m- 1) + (n - 1) ; line++ {
		arr := listMap[line]
		if line % 2 == 0 {
			for i := len(arr) - 1; i > -1; i-- {
			 	res = append(res, arr[i])
			}
		} else {
			for i := 0; i < len(arr); i++ {
				res = append(res, arr[i])
			}
		}
	}

	return res
}






