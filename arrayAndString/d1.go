package arrayAndString

import (
	"sort"
)

//type IntArray [][]int
//
//func (s IntArray) len() int {
//	return len(s)
//}
//
////func (s IntArray) Less(i, j int) bool {
////	return s[i][0] > s[j][0]
////}
//
//func (s IntArray) Less(i, j int) bool {
//	return s[i][0] > s[j][0]
//}
//
//func (s IntArray) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] > intervals[j][0] {
			return false
		}
		return true
	})

	res := [][]int{}
	temp := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if temp[0] <=intervals[i][0] && intervals[i][0] <= temp[1] {
			temp[1] = intervals[i][1]
		} else {
			res = append(res, temp)
			temp = intervals[i]
		}
	}
	res = append(res, temp)
	return res
}