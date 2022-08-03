package dp

// 超时
func wordBreak(s string, wordDict []string) bool {
	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}

	// 记录重复递归的结果
	dfsMap := make(map[string]bool)

	var dfs func(leftS string) bool
	dfs = func(leftS string) bool {
		if r, ok := dfsMap[leftS]; ok {
			return r
		}
		if leftS == "" {
			return true
		}
		res := false
		for i := 1; i <= len(leftS); i++ {
			if wordMap[leftS[:i]] {
				res = dfs(leftS[i:]) || res
			}
		}
		dfsMap[leftS] = res
		return res
	}

	return dfs(s)
}

func wordBreak2(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true

	wordMap := make(map[string]bool)
	for _, w := range wordDict {
		wordMap[w] = true
	}

	for j := 1; j <= len(s); j++ {
		for i := 0; i < j; i++ {
			if dp[i] && wordMap[s[i:j]] {
				dp[j] = true
				break
			}
		}
	}

	return dp[len(s)]
}
