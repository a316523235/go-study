package str

import (
	"math"
	"strconv"
	"strings"
)

func myAtoi(s string) int {
	s = strings.TrimSpace(s)
	bt := []byte(s)
	rev, isStart, isFu := 0, false, false

	for i := 0; i < len(bt); i++ {
		switch {
		case bt[i] == '+' && !isStart:
			isStart = true
			isFu = false
		case bt[i] == '-' && !isStart:
			isStart = true
			isFu = true
		case isNum(bt[i]):
			isStart = true
			n, _ := strconv.Atoi(string(bt[i]))
			rev = rev * 10 + n
			if rev > math.MaxInt32 {
				goto over
			}
		default:
			goto over
		}
	}
	over:
		switch {
		case rev == 0:
			return 0
		case isFu:
			rev = rev * -1
			if rev < math.MinInt32 {
				return  math.MinInt32
			}
			return rev
		default:
			if rev > math.MaxInt32 {
				return math.MaxInt32
			}
			return rev
		}
}

func isNum(s byte) bool {
	return s >= '0' && s <= '9'
}