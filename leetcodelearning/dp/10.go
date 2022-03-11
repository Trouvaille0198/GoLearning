package dp

func isMatch(s string, p string) bool {
	row, col := len(s), len(p)
	match := func(i, j int) bool {
		if i == 0 {
			return false
		}
		if p[j-1] == '.' {
			return true
		}
		return s[i-1] == p[j-1]
	}

	dp := make([][]bool, row+1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]bool, col+1)
	}
	dp[0][0] = true

	for i := 2; i <= col; i++ {
		if p[i-1] == '*' {
			dp[0][i] = dp[0][i-2]
		}
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if p[j-1] != '*' {
				if match(i, j) {
					dp[i][j] = dp[i-1][j-1]
				} else {
					dp[i][j] = false
				}
			} else {
				if match(i, j-1) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[row][col]
}
