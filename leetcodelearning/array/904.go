package array

// DP
func totalFruit(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	dp := make([]int, len(fruits))
	dp[0] = 1
	a, b, aIndex, bIndex := -1, fruits[0], -1, 0
	maxResult := 1
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == a || fruits[i] == b {
			// 加老果子
			if fruits[i] == a {
				aIndex = i
			}
			if fruits[i] == b {
				bIndex = i
			}
			dp[i] = dp[i-1] + 1
		} else {
			// 换新果子
			var length int
			// 计算最近一类果实的长度 并覆盖掉丢弃果子的种类
			if aIndex > bIndex {
				length = aIndex - bIndex
				b = fruits[i]
				bIndex = i
			} else {
				length = bIndex - aIndex
				a = fruits[i]
				aIndex = i
			}
			dp[i] = length + 1
		}
		if maxResult < dp[i] {
			maxResult = dp[i]
		}
	}
	return maxResult
}

// DP better
func totalFruit2(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	a, b, aIndex, bIndex := -1, fruits[0], -1, 0
	pre, curResult, maxResult := 1, 1, 1 // 记录dp[i-1] dp[i] max(dp[i])
	for i := 1; i < len(fruits); i++ {
		if fruits[i] == a || fruits[i] == b {
			// 加老果子
			if fruits[i] == a {
				aIndex = i
			}
			if fruits[i] == b {
				bIndex = i
			}
			curResult = pre + 1
		} else {
			// 换新果子
			var length int
			// 计算最近一类果实的长度 并覆盖掉丢弃果子的种类
			if aIndex > bIndex {
				length = aIndex - bIndex
				b = fruits[i]
				bIndex = i
			} else {
				length = bIndex - aIndex
				a = fruits[i]
				aIndex = i
			}
			curResult = length + 1
		}
		pre = curResult
		if maxResult < curResult {
			maxResult = curResult
		}
	}
	return maxResult
}

// 滑动窗口
func totalFruit3(fruits []int) int {
	if len(fruits) == 0 {
		return 0
	}
	a, b, aIndex, bIndex := -1, fruits[0], -1, 0
	i, j := -1, 0 // i为左边界 保证i+1到j
	maxLen := 0
	for j < len(fruits) {
		if fruits[j] != a && fruits[j] != b {
			// j是新果子
			if aIndex > bIndex {
				// 将i移至淘汰的旧果子最后的位置
				i = bIndex
				// 替换旧果子
				b = fruits[j]
				bIndex = j
			} else {
				i = aIndex
				a = fruits[j]
				aIndex = j
			}
		} else {
			// j是老果子
			if fruits[j] == a {
				aIndex = j
			}
			if fruits[j] == b {
				bIndex = j
			}
		}
		if j-i > maxLen {
			maxLen = j - i
		}
		j++
	}
	return maxLen
}
