package dp

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	hash := make(map[byte]int) // 记录每个字母最大的索引
	dp := make([]int, len(s))
	dp[0] = 1
	hash[s[0]] = 0

	maxResult := 1
	for i := 1; i < len(s); i++ {
		index, ok := hash[s[i]]
		if !ok {
			// s[i]未出现过
			dp[i] = dp[i-1] + 1
		} else {
			if i-index > dp[i-1] {
				// 上一个与s[i]相等的字母下标在dp[i-1]之前 可以加1
				dp[i] = dp[i-1] + 1
			} else {
				// 上一个与s[i]相等的字母下标在dp[i-1]之中 没有办法
				dp[i] = i - index
			}
		}
		hash[s[i]] = i
		if maxResult < dp[i] {
			maxResult = dp[i]
		}
	}
	return maxResult
}

func lengthOfLongestSubstring2(s string) int {
	if len(s) == 0 {
		return 0
	}
	hash := make(map[byte]int) // 记录每个字母最大的索引

	hash[s[0]] = 0
	maxResult := 1
	pre := 1          // 记录dp[i-1]
	var curResult int // 记录dp[i]

	for i := 1; i < len(s); i++ {
		index, ok := hash[s[i]]
		if !ok {
			// s[i]未出现过
			curResult = pre + 1
		} else {
			if i-index > pre {
				// 上一个与s[i]相等的字母下标在dp[i-1]之前 可以加1
				curResult = pre + 1
			} else {
				// 上一个与s[i]相等的字母下标在dp[i-1]之中 没有办法
				curResult = i - index
			}
		}
		hash[s[i]] = i // 更新
		pre = curResult
		if maxResult < curResult {
			maxResult = curResult
		}
	}
	return maxResult
}

// 双指针
func lengthOfLongestSubstring3(s string) int {
	if len(s) == 0 {
		return 0
	}
	hash := make(map[byte]int) // 记录每个字母最大的索引
	maxResult := 1

	i := -1 // 左指针 定义左边界 保证i+1到j的字串无重复
	for j := 0; j < len(s); j++ {
		lastIndex, ok := hash[s[j]]

		if ok && lastIndex > i {
			// 上一个与s[j]相等的字母下标在左边界i右边 以其为新左边界
			i = lastIndex
		}
		hash[s[j]] = j // 更新

		if maxResult < j-i {
			maxResult = j - i
		}
	}
	return maxResult
}
