package string

// Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.
func romanToInt(s string) int {
	roman := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && roman[string(s[i])] < roman[string(s[i+1])] {
			result -= roman[string(s[i])]
		} else {
			result += roman[string(s[i])]
		}
	}
	return result
}
