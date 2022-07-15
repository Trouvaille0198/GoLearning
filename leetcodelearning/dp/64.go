package leetcodelearning

import "math"

func minPathSum(grid [][]int) int {
	res := math.MaxInt32
	m, n := len(grid), len(grid[0])
	var dfs func(x, y, sum int)
	dfs = func(x, y, sum int) {
		if x >= m || y >= n {
			return
		}
		sum += grid[x][y]
		if x == m-1 && y == n-1 {
			if sum <= res {
				res = sum
			}
			return
		}
		if sum >= res {
			return
		}
		dfs(x+1, y, sum)
		dfs(x, y+1, sum)
	}
	dfs(0, 0, 0)
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = grid[0][0]
	for x := 1; x < m; x++ {
		dp[x][0] = dp[x-1][0] + grid[x][0]
	}
	for y := 1; y < n; y++ {
		dp[0][y] = dp[0][y-1] + grid[0][y]
	}

	for x := 1; x < m; x++ {
		for y := 1; y < n; y++ {
			dp[x][y] = min(dp[x-1][y], dp[x][y-1]) + grid[x][y]
		}
	}
	return dp[m-1][n-1]
}
