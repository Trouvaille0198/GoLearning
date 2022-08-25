package dp

func maximalSquare(matrix [][]byte) int {
	rows, cols := len(matrix), len(matrix[0])
	dp := make([][]int, rows)
	for i := 0; i < rows; i++ {
		dp[i] = make([]int, cols)
	}
	maxSideLen := 0
	// 先给第0行和第0列赋值
	for y := 0; y < cols; y++ {
		if matrix[0][y] == '1' {
			dp[0][y] = 1
			maxSideLen = 1
		} else {
			dp[0][y] = 0
		}
	}
	for x := 0; x < rows; x++ {
		if matrix[x][0] == '1' {
			dp[x][0] = 1
			maxSideLen = 1
		} else {
			dp[x][0] = 0
		}
	}
	for x := 1; x < rows; x++ {
		for y := 1; y < cols; y++ {
			if matrix[x][y] == '1' {
				dp[x][y] = min(dp[x-1][y-1], min(dp[x-1][y], dp[x][y-1])) + 1
				if dp[x][y] > maxSideLen {
					maxSideLen = dp[x][y]
				}
			}
		}
	}
	return maxSideLen * maxSideLen
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
