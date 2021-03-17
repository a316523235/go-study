package simple

import (
	"sort"
)

func rotate(nums []int, k int) {
	swapCount := 0
	for timeIndex := 0; swapCount < len(nums); timeIndex++ {
		preIndex := timeIndex
		pre := nums[timeIndex]
		for {
			next := (preIndex + k) % len(nums)
			nums[next], pre = pre, nums[next]
			preIndex = next
			swapCount++
			if next == timeIndex {
				break
			}
		}
	}
}

func containsDuplicate(nums []int) bool {
	mp := map[int]bool{}
	for _, n := range nums {
		if mp[n] {
			return true
		}
		mp[n] = true
	}
	return false
}

func intersect(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	var res []int
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] < nums2[j] {
			i++
		} else {
			j++
		}
	}
	return res
}

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += 1
		if digits[i] < 10 {
			break
		} else {
			digits[i] = 0
			if i == 0 {
				digits = append([]int{1}, digits...)
			}
		}
	}
	return digits
}

func isValidSudoku(board [][]byte) bool {
	arrRow, arrCol, arr9 := [9][9]bool{}, [9][9]bool{}, [9][9]bool{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			num := board[i][j] - byte('1')
			k := (i/3)*3 + j/3
			if arrRow[i][num] || arrCol[j][num] || arr9[k][num] {
				return false
			}
			arrRow[i][num], arrCol[j][num], arr9[k][num] = true, true, true
		}
	}
	return true
}

func rotate2(matrix [][]int) {
	n := len(matrix[0])
	groupCnt := n / 2

	for groupIndex := 0; groupIndex < groupCnt; groupIndex++ {
		for i := groupIndex; i < n-1-groupIndex; i++ {
			i1, j1 := groupIndex, i          //左上
			i2, j2 := j1, n-1-groupIndex     //右上
			i3, j3 := n-1-groupIndex, n-1-i2 //右下
			i4, j4 := j3, groupIndex         //左下

			startNum := matrix[i1][j1]

			matrix[i1][j1] = matrix[i4][j4]
			matrix[i4][j4] = matrix[i3][j3]
			matrix[i3][j3] = matrix[i2][j2]
			matrix[i2][j2] = startNum
		}
	}
}
