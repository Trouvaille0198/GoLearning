package back_tracking

func generateParenthesis(n int) []string {
	count := n * 2
	res := make([]string, 0)
	var dfs func(quotes []byte, layer, leftNum int) // leftNum 表示之前剩余多少个没有被匹配到的左括号
	dfs = func(quotes []byte, layer, leftNum int) {
		if layer == count {
			if leftNum == 0 {
				// 这里不用 copy, 因为 string 函数直接深拷贝了
				res = append(res, string(quotes))
			}
			return
		}

		if leftNum >= 1 {
			quotes = append(quotes, ')')
			dfs(quotes, layer+1, leftNum-1)
			quotes = quotes[:len(quotes)-1]
		}
		quotes = append(quotes, '(')
		dfs(quotes, layer+1, leftNum+1)
		quotes = quotes[:len(quotes)-1]
	}
	dfs([]byte{}, 0, 0)
	return res
}

func generateParenthesis(n int) []string {
	if n == 1 {
		return []string{"()"}
	}
	dp := [][]string{{""}, {"()"}}

	for i := 2; i <= n; i++ {
		dpi := make([]string, 0) // dp[i]
		for p := 0; p <= i-1; p++ {
			q := i - 1 - p
			// 遍历dp[p]和dp[q]中的所有组合
			for _, quotes1 := range dp[p] {
				for _, quotes2 := range dp[q] {
					dpi = append(dpi, "("+quotes1+")"+quotes2)
				}
			}
		}
		dp = append(dp, dpi)
	}
	return dp[n]
}
