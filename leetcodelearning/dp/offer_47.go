package dp

func maxValue(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	dp := make([][]int, 0)
	for i := 0; i < m; i++ {
		tmp := make([]int, n)
		dp = append(dp, tmp)
	}

	var up, left int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
				continue
			}
			if i > 0 {
				up = dp[i-1][j] + grid[i][j]
			} else {
				up = 0
			}
			if j > 0 {
				left = dp[i][j-1] + grid[i][j]
			} else {
				left = 0
			}

			if up > left {
				dp[i][j] = up
			} else {
				dp[i][j] = left
			}
		}
	}
	return dp[m-1][n-1]
}

func maxValue2(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	var up, left int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 {
				up = grid[i-1][j] + grid[i][j]
			} else {
				up = 0
			}
			if j > 0 {
				left = grid[i][j-1] + grid[i][j]
			} else {
				left = 0
			}

			if up > left {
				grid[i][j] = up
			} else {
				grid[i][j] = left
			}
		}
	}
	return grid[m-1][n-1]
}

func maxValue3(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	for i := 1; i < n; i++ {
		grid[0][i] += grid[0][i-1]
	}
	for j := 1; j < m; j++ {
		grid[j][0] += grid[j-1][0]
	}

	var up, left int
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			up = grid[i-1][j] + grid[i][j]
			left = grid[i][j-1] + grid[i][j]

			if up > left {
				grid[i][j] = up
			} else {
				grid[i][j] = left
			}
		}
	}
	return grid[m-1][n-1]
}
