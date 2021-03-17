package computer

func countPrimes(n int) int {
	cnt := 0
	dp := make([]int, n + 1)
	for number := 2; number <= n; number++ {
		dp[number] = 1
	}
	for number := 2; number < n; number++ {
		if dp[number] == 1 {
			for i := 2; i*number <= n; i++ {
				dp[i*number] = 0
			}
		}
	}
	return cnt
}

func isPrimes(n int) bool {
	return true
}
