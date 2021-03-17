package str

import (
	"math"
	"sort"
	"strings"
)

func reverseString(s []byte)  {
	for i, j := 0, len(s); i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func reverse(x int) int {
	rev := 0
	for x != 0 {
		pop := x % 10
		x /= 10
		if rev > math.MaxInt32 / 10 || (rev == math.MaxInt32 / 10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32 / 10 || (rev == math.MinInt32 / 10 && pop < -8) {
			return 0
		}
		rev = rev*10 + pop
	}
	return rev
}

func firstUniqChar(s string) int {
	mp := map[rune]int{}
	for _, c := range s {
		mp[c]++
	}
	for i, c := range s {
		if mp[c] == 1 {
			return i
		}
	}
	return -1
}

func isAnagram(s string, t string) bool {
	bt1, bt2 := []byte(s), []byte(t)
	sort.Slice(bt1, func(i, j int) bool {
		return bt1[i] < bt1[j]
	})
	sort.Slice(bt2, func(i, j int) bool {
		return bt2[i] < bt2[j]
	})

	s, t = string(bt1), string(bt2)

	if s == t {
		return true
	}
	return false
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	bt := []byte(s)
	for ; i < j; {
		if !isalnum(bt[i]) {
			i++
			continue
		}
		if !isalnum(bt[j]) {
			j--
			continue
		}
		if bt[i] != bt[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func isalnum(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9')
}