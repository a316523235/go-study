package computer

func romanToInt(s string) int {
	res := 0
	pre := 0
	for i := len(s); i > -1; i-- {
		temp := which(s[i])
		if temp < pre {
			res -= temp
		} else {
			res += temp
		}
		pre = temp
	}
	return res
}

func which(ch uint8) int {
	switch ch {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}