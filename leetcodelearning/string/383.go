package string

func canConstruct(ransomNote string, magazine string) bool {
	hash := make(map[byte]int)
	for _, letter := range []byte(magazine) {
		hash[letter]++
	}
	for _, letter := range []byte(ransomNote) {
		hash[letter]--
		if hash[letter] < 0 {
			return false
		}
	}
	return true
}

func canConstruct(ransomNote string, magazine string) bool {
	var hash [26]int
	for _, letter := range []byte(magazine) {
		hash[letter-'a']++
	}
	for _, letter := range []byte(ransomNote) {
		hash[letter-'a']--
		if hash[letter-'a'] < 0 {
			return false
		}
	}
	return true
}
