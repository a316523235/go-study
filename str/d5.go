package str

import "strconv"

func strStr(haystack string, needle string) int {
	len1, len2 := len(haystack), len(needle)
	if len2 > len1 {
		return -1
	}
	sum1, sum2 := 0, 0
	for i := 0; i < len2; i++ {
		sum1 += int(haystack[i])
	}
	for i := 0; i < len2; i++ {
		sum2 += int(needle[i])
	}
	if haystack[0:len2] == needle[0:len2] {
		return 0
	}
	for i := len2; i < len1; i++ {
		sum1 = sum1 - int(haystack[i-len2])  + int(haystack[i])
		if sum1 == sum2 && haystack[i-len2 + 1: i+1] == needle {
			return i - len2 + 1
		}
	}
	return -1
}

func countAndSay(n int) string {
	str := "1"
	for i := 1; i < n; i++ {
		newStr := ""
		startIndex := 0
		for j := 1; j <= len(str); j++ {
			isEqual := j < len(str) && str[j] == str[startIndex]
			if isEqual {
				continue
			}
			newStr = newStr + strconv.Itoa(j - startIndex) + string(str[startIndex])
			startIndex = j
		}
		str = newStr
	}
	return str
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	i := 0
	for {
		for arrIndex := 0; arrIndex < len(strs); arrIndex++ {
			if i >= len(strs[arrIndex]) || strs[arrIndex][i] != strs[0][i] {
				return strs[0][:i]
			}
		}
		i++
	}
}
