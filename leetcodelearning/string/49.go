package string

func groupAnagrams(strs []string) [][]string {
	hash := make(map[[26]int32][]string)
	for _, str := range strs {
		var count [26]int32
		for _, letter := range str {
			count[letter-'a']++
		}
		hash[count] = append(hash[count], str)
	}
	res := make([][]string, 0, len(hash))
	for _, v := range hash {
		res = append(res, v)
	}
	return res
}
