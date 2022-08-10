package dp

import "math"

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minCnt := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minCnt = min(minCnt, dp[i-j*j])
		}
		dp[i] = minCnt + 1
	}
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
