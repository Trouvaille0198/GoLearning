package array

import (
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	i, j := 0, 0
	minLen := math.MaxInt
	var curSum int
	for j < len(nums) {
		curSum += nums[j] // 加上新边界
		for curSum >= target {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			curSum -= nums[i] // 减去旧的左边界
			i++
		}
		j++

	}

	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}
