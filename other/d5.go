package other

func isValid(s string) bool {
	n := len(s)
	if n % 2 != 0 {
		return false
	}
	mp := map[byte]byte{
		'}':'{',
		']':'[',
		')':'(',
	}
	pre := []byte{}
	for i := 0; i < n; i++ {
		if len(pre) == 0 {
			pre = append(pre, s[i])
			continue
			//mp[左括号] 结果为空， 不会与前一个字符相等
		} else if pre[len(pre) - 1] == mp[s[i]] {
			pre = pre[:len(pre) - 1]
		} else {
			pre = append(pre, s[i])
		}
	}
	return len(pre) == 0
}

