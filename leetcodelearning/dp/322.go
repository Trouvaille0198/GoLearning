package dp

import "math"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}

	for i := 0; i <= amount; i++ {
		minVal := math.MaxInt32
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] != -1 && dp[i-coin]+1 < minVal {
				// 硬币面额小于等于i && dp[找零]存在 && dp[找零]+1暂时最小
				minVal = dp[i-coin] + 1
			}
		}
		if minVal != math.MaxInt32 {
			dp[i] = minVal
		}
	}
	return dp[amount]
}
