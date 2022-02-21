package string

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	hash := make(map[byte]bool) // 记录每个字符是否出现过
	var left, right, maxLen int = 0, 0, 0
	for right < len(s) {
		for hash[s[right]] {
			// 出现重复 左边界一直右移 直到不重复
			hash[s[left]] = false
			left++
		}
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
		hash[s[right]] = true
		right++
	}
	return maxLen
}
