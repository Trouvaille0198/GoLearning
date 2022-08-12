package dp

func maxProfit(prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(prices); i++ {
		dp[i] = make([]int, 3)
	}

	dp[0] = []int{-prices[0], 0, 0}

	for i := 1; i < len(prices); i++ {
		// 持有股票
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i]) // 之前的/冷冻期后新买的
		// 不持有股票
		dp[i][1] = max(dp[i-1][0]+prices[i], dp[i-1][1]) // 今天卖掉了/之前就没有
		// 冷冻期
		dp[i][2] = dp[i-1][1] // 昨天卖掉的
	}
	return dp[len(prices)-1][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
