package other

func hammingDistance(x int, y int) int {
	xor := x ^ y
	cnt  := 0
	//mask == 0时表示位移已经超过32位，再移已经没意义
	for mask := 1; xor >= mask && mask != 0; mask <<= 1 {
		if xor & mask != 0 {
			cnt++
		}
	}
	return cnt
}
