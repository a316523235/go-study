package computer

import "math"

var dpPower map[int]bool

func isPowerOfThree(n int) bool {
	initMap()
	return dpPower[n]
}

func initMap()  {
	if dpPower != nil {
		return
	}
	dpPower = map[int]bool{}
	//3的0次方是1, 所以从1开始
	for n := 1; n < math.MaxInt32; n = n * 3 {
		dpPower[n] = true
	}
}