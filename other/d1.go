package other

func hammingWeight(num uint32) int {
	var mask uint32
	cnt  := 0
	mask = 1
	for maskCnt := 0; maskCnt < 31 && num > mask; maskCnt++ {
		if num & mask != 0 {
			cnt++
		}
		mask <<= 1
		//fmt.Println(mask)
		//time.Sleep(1 * time.Second)
	}
	//fmt.Println(maskCnt)
	return cnt
}