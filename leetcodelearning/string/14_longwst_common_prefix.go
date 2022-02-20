package string

func longestCommonPrefix(strs []string) string {
	prefix := strs[0] // 先随便找一个作为前缀 在遍历过程中再慢慢剔除

	for i := 1; i < len(strs); i++ {
		// 遍历每个字符串
		for j := 0; j < len(prefix); j++ {
			// 遍历每个前缀字母
			if len(strs[i]) <= j || strs[i][j] != prefix[j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}
