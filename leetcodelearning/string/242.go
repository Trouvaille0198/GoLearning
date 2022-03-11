package string

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	hash := make(map[byte]int)
	for _, letter := range s {
		hash[byte(letter)]++
	}
	for _, letter := range t {
		fmt.Println(letter)
		if count, ok := hash[byte(letter)]; !ok || count == 0 {
			// 没有或已经减完
			return false
		}
		hash[byte(letter)]--
	}
	for _, v := range hash {
		if v != 0 {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var hash [26]int
	for _, letter := range s {
		hash[letter-'a']++
	}
	for _, letter := range t {
		if hash[letter-'a'] == 0 {
			return false
		}
		hash[letter-'a']--
	}
	for i := 0; i < 26; i++ {
		if hash[i] > 0 {
			return false
		}
	}
	return true
	// 甚至可以酱紫写 👇
	// return record == [26]int{}
}
