package simple

func maxProfit(prices []int) int {
	var dp [][]int
	for i := 0; i < len(prices); i++ {
		dp = append(dp, []int{0, 0})
	}
	dp[0][0] = 0	//未持有
	dp[0][1] = -prices[0] //持有

	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}

	return dp[len(prices) - 1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
