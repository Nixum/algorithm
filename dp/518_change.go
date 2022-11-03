package dp

// 求数量：dp[j] += dp[j - coins[i]]
// 如果求组合数就是外层for循环遍历物品，内层for遍历背包。
// 如果求排列数就是外层for遍历背包，内层for循环遍历物品。
func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]]
		}
	}
	return dp[amount]
}
