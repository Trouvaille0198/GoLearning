package string

// strStr KMP算法 时间复杂度O(N+M),空间复杂度O(M)
func strStrSimple(haystack string, needle string) int {
	n, m := len(haystack), len(needle)
	if m == 0 {
		return 0
	}
	next := make([]int, m)
	GetNext(next, needle)
	// 因为next数组里记录的起始位置为0
	j := 0
	// i从0开始匹配
	for i := 0; i < n; i++ {
		// 如果不匹配，就寻找之前匹配的位置
		for j > 0 && haystack[i] != needle[j] {
			j = next[j-1]
		}
		// 如果匹配，i和j同时向后移动
		if haystack[i] == needle[j] {
			j++
		}
		// 如果j从0移动到m的位置，意味着模式串needle与文本串haystack匹配成功
		if j == m {
			return i - m + 1
		}
	}
	return -1
}

func GetNext(next []int, s string) {
	// next[j]就是记录着j（包括j）之前的子串的相同前后缀的长度。
	j := 0
	next[0] = 0
	// j指向前缀起始位置，i指向后缀起始位置
	for i := 1; i < len(s); i++ {
		// 如果前后缀不相同，那么j就要向前回退
		for j > 0 && s[i] != s[j] {
			j = next[j-1]
		}
		// 说明找到了相同的前后缀, j++，同时记录next[i]
		if s[i] == s[j] {
			j++
		}
		``
		next[i] = j
	}
}
