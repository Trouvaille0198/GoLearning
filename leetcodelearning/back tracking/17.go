package back_tracking

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	res := make([]string, 0)
	var bt func(index int, combination string)
	bt = func(index int, combination string) {
		if index == len(digits) {
			res = append(res, combination)
			// tmpStr = ""
			return
		}

		for _, word := range phoneMap[string(digits[index])] {
			bt(index+1, combination+string(word))
		}
	}
	bt(0, "")
	return res
}
