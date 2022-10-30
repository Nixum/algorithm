package dp

// 时间复杂度和空间复杂度都是O(n)
func fib(n int) int {
	if n <= 1 {
		return n
	}
	fb := make([]int, n + 1)
	fb[0] = 0
	fb[1] = 1
	for i := 2; i <= n; i++ {
		fb[i] = fb[i - 1] + fb[i - 2]
	}
	return fb[n]
}

// 优化，时间复杂度O(n)，空间复杂度O(1)
func fib2(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, 2)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
