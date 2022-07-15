package string

func repeatedSubstringPattern(s string) bool {
	for i := 1; i <= len(s)/2; i++ {
		// i代表长度
		if len(s)%i == 0 {
			// 长度整除
			isMatch := true
			for j := i; j < len(s); j++ {
				// 判断之后的字符串是否可以由0到i-1的子串构成
				if s[j] != s[j-i] {
					isMatch = false
					break
				}
			}
			if isMatch {
				return true
			}
		}
	}
	return false
}

func getNext(s string) []int {
	next := make([]int, len(s)) // 0-i子串的最长公共前后缀
	next[0] = 0
	i := 0 // i指向前缀
	for j := 1; j < len(s); j++ {
		// j指向后缀
		for i > 0 && s[i] != s[j] {
			i = next[i-1]
		}
		if s[i] == s[j] {
			i++
		}
		next[j] = i
	}
	return next
}

func repeatedSubstringPattern(s string) bool {
	if len(s) == 0 {
		return false
	}
	next := getNext(s)
	if next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	}
	return false
}
