package other

func reverseBits(num uint32) uint32 {
	var res uint32
	for i := 0; i < 32; i++ {
		bit := num & 1
		res <<= 1
		res += bit
		num >>= 1
	}
	return res
}

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i <= numRows; i++ {
		arr := make([]int, i + 1)
		for j := 0; j <= i; j++ {
			if i < 2 || j == 0 || j == i - 1 {
				arr[j] = 1
			} else {
				arr[j] = res[i - 1][j - 1] + res[i-1][j]
			}
		}
		res = append(res, arr)
	}
	return res
}
