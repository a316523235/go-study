package simple

import (
	"fmt"
	"math"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	prices := []int{7,1,5,3,6,4}
	res := maxProfit(prices)
	fmt.Println(res)
}

func TestIsValidSudoku(t *testing.T) {
	source := [][]byte{
		{'.','.','4',   '.','.','.',   '6','3','.'},
		{'.','.','.',   '.','.','.',   '.','.','.'},
		{'5','.','.',   '.','.','.',   '.','9','.'},

		{'.','.','.',   '5','6','.',   '.','.','.'},
		{'4','.','3',   '.','.','.',   '.','.','1'},
		{'.','.','.',   '7','.','.',   '.','.','.'},

		{'.','.','.',   '5','.','.',   '.','.','.'},
		{'.','.','.',   '.','.','.',   '.','.','.'},
		{'.','.','.',   '.','.','.',   '.','.','.'}}
	res := isValidSudoku(source)
	fmt.Println(res)
}

func TestIsValidSudoku2(t *testing.T) {
	arr := []int{}
	arr[3] = 3
	fmt.Println(arr)
}

//核心重点 dp[i] = max(dp[i-1] + nums[i], nums[i])
func maxSubArray(nums []int) int {
	return maxSubArrayByDp(nums)
	maxSub, currentSub := math.MinInt32, math.MinInt32
	for i := 0; i < len(nums); i++ {
		currentSub = max(nums[i], currentSub + nums[i])
		maxSub = max(maxSub, currentSub)
	}
	return maxSub
}

//核心重点 dp[i] = max(dp[i-1] + nums[i], nums[i])
func maxSubArrayByDp(nums []int) int {
	dp := []int{nums[0]}
	maxSub := nums[0]
	for i := 1; i < len(nums); i++ {
		dp = append(dp, 0)
		dp[i] = max(nums[i], dp[0] + nums[i])
		maxSub = max(maxSub, dp[i])
	}
	return maxSub
}