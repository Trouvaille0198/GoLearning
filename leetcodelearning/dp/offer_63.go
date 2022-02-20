package dp

import "math"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 一次遍历
func maxProfit(prices []int) int {
	curMinPrice := math.MaxInt
	maxPro := 1

	for i := 0; i < len(prices); i++ {
		maxPro = max(maxPro, prices[i]-curMinPrice)
		curMinPrice = min(curMinPrice, prices[i])
	}
	return maxPro
}

// DP
func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	dp := make([]int, len(prices))
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if dp[i-1] >= prices[i]-minPrice {
			dp[i] = dp[i-1]
		} else {
			dp[i] = prices[i] - minPrice
		}

		if prices[i] < minPrice {
			minPrice = prices[i]
		}
	}
	return dp[len(prices)-1]
}
