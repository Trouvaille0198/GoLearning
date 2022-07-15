package dp

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i/2; j++ {
			// 遍历每一个结点作为头节点 因为是对称的情况所以可以遍历一半
			dp[i] += dp[j-1] * dp[i-j]
		}
		dp[i] *= 2
		if i%2 == 1 {
			// 单独处理奇数的情况
			dp[i] += dp[i/2] * dp[i/2]
		}
	}
	return dp[n]
}

// 易于理解的递归思路
func numTrees(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	sum := 0
	for i := 1; i <= n; i++ {
		sum += numTrees(i-1) * numTrees(n-i)
	}
	return sum
}
