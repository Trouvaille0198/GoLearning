package dp

func maxProduct(nums []int) int {
	maxRes, minRes, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		maxTmp, minTmp := maxRes, minRes
		maxRes = max(maxTmp*nums[i], max(minTmp*nums[i], nums[i]))
		minRes = min(minTmp*nums[i], min(maxTmp*nums[i], nums[i]))
		res = max(maxRes, res)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
