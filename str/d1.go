package str

func firstUniqChar2(s string) int {
	mp := [255]int{}
	for _, v := range s {
		mp[v]++
	}
	for i, v := range s{
		if mp[v] == 1 {
			return i
		}
	}
	return -1
}
