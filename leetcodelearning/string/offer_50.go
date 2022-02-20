package string

import "math"

func firstUniqChar(s string) byte {
	hash := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		hash[s[i]] = append(hash[s[i]], i)
	}
	resultWord := byte(' ')
	index := math.MaxInt
	for k, v := range hash {
		if len(v) == 1 && v[len(v)-1] < index {
			resultWord = k
			index = v[len(v)-1]
		}
	}
	return resultWord
}

func firstUniqChar2(s string) byte {
	var list [26]int
	length := len(s)
	for i := 0; i < length; i++ {
		list[s[i]-'a']++
	}
	for i := 0; i < length; i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
