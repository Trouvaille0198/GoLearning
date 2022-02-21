package string

import (
	"math"
)

func minWindow(s string, t string) string {
	if s == t {
		return s
	}
	hash := make(map[byte]int) // 记录目标字母及其个数 作为需求
	for i := 0; i < len(t); i++ {
		hash[t[i]]++
	}
	reL, reR := 0, -1 // 目标的左右边界
	minLen := math.MaxInt
	left, right := 0, 0
	sum := len(t) // 记录总数 为0时代表全命中

	for right < len(s) {
		if _, ok := hash[s[right]]; ok {
			// 存在
			hash[s[right]]--
			if hash[s[right]] >= 0 {
				// 变成负数时sum不用减少 因为已经满足
				sum--
			}
		}

		if sum == 0 {
			// 全部命中
			for left <= right {
				// left右移 直到不满足全命中
				if _, ok := hash[s[left]]; ok {
					hash[s[left]]++
					if hash[s[left]] > 0 {
						// 需求得不到满足了
						if minLen > right-left+1 {
							// 刷新最短值
							reR, reL = right, left
							minLen = right - left + 1
						}
						sum++
						left++ // 正式地++ 破坏全命中状态
						break
					}
				}
				left++ // 注意 先判断再++
			}
		}
		right++
	}
	return s[reL : reR+1]
}
