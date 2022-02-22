package string

import "strings"

func reverseWords(s string) string {
	words := make([]string, 0)
	i, j := len(s)-1, len(s)-1
	for i >= 0 && j >= 0 {
		// j先移动到一个单词的右边界
		for j >= 0 && s[j] == ' ' {
			j--
		}
		if j < 0 {
			break
		}
		i = j
		// i移动到一个单词的左边界
		for i >= 0 && s[i] != ' ' {
			i--
		}
		words = append(words, s[i+1:j+1])
		j = i
	}
	return strings.Join(words, " ")
}
