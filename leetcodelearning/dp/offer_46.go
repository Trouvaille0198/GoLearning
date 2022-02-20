package dp

import (
	"strconv"
)

func translateNum(num int) int {
	s := strconv.Itoa(num)
	if len(s) == 0 {
		return 0
	}
	dp := make([]int, len(s))
	dp[0] = 1
	if len(s) == 1 {
		return 1
	}
	if s[0:2] <= "25" && s[0:2] >= "10" {
		dp[1] = 2
	} else {
		dp[1] = 1
	}

	for i := 2; i < len(s); i++ {
		pre := s[i-1 : i+1]
		if pre <= "25" && pre >= "10" {
			dp[i] = dp[i-1] + dp[i-2]
		} else {
			dp[i] = dp[i-1]
		}
	}
	return dp[len(s)-1]
}

func translateNum2(num int) int {
	s := strconv.Itoa(num)
	if len(s) == 0 {
		return 0
	}
	a1, a2 := 1, 1 // dp[i-1]和dp[i-2]
	result := 1
	for i := 1; i < len(s); i++ {
		pre := s[i-1 : i+1] // 当前位和上一位组成的二位数字
		if pre <= "25" && pre >= "10" {
			result = a1 + a2
		} else {
			result = a1
		}
		a2 = a1
		a1 = result
	}
	return result
}
