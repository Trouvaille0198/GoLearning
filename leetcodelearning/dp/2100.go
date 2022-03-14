package dp

func goodDaysToRobBank(security []int, time int) []int {
	res := make([]int, 0)
	dp1 := make([]int, len(security))
	dp1[0] = 0
	for i := 1; i < len(security); i++ {
		if security[i-1] >= security[i] {
			dp1[i] = dp1[i-1] + 1
		} else {
			dp1[i] = 0
		}
	}
	dp2 := make([]int, len(security))
	dp2[len(security)] = 0
	for i := len(security) - 1; i >= 0; i-- {
		if security[i] <= security[i+1] {
			dp2[i] = dp2[i+1] + 1
		} else {
			dp2[i] = 0
		}
	}

	for i := time; i < len(security)-time; i++ {
		if dp1[i] >= time && dp2[i] >= time {
			res = append(res, i)
		}
	}
	return res
}
