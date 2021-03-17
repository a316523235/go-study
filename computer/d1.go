package computer

import "strconv"

/*
写一个程序，输出从 1 到 n 数字的字符串表示。

1. 如果 n 是3的倍数，输出“Fizz”；

2. 如果 n 是5的倍数，输出“Buzz”；

3.如果 n 同时是3和5的倍数，输出 “FizzBuzz”。
 */

var nums = []int{3, 5}
var texts = []string{"Fizz", "Buzz"}

func fizzBuzz(n int) []string {
	res := []string{}
	for i := 0; i <= n; i++ {
		ansStr := ""
		for j := 0; j < len(nums); j++ {
			if i % nums[j] == 0 {
				ansStr += texts[j]
			}
		}
		if ansStr == "" {
			ansStr = strconv.Itoa(i)
		}
		res = append(res, ansStr)
	}
	return res
}