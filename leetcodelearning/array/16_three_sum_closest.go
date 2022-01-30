package array

import (
	"math"
	"sort"
)

// Given an array nums of n integers and an integer target,
// find three integers in nums such that the sum is closest to target. Return the sum of the three integers.
// You may assume that each input would have exactly one solution.

// Given array nums = [-1, 2, 1, -4], and target = 1.
// The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// threeSumClosest 排序+双指针
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt32
	resultSum := math.MaxInt32
	for index := 1; index < len(nums)-1; index++ {
		start, end := 0, len(nums)-1
		for start < index && end > index {
			curSum := nums[start] + nums[index] + nums[end]
			curDiff := curSum - target

			if abs(curDiff) < minDiff {
				minDiff = abs(curDiff)
				resultSum = curSum
			}

			switch {
			case curDiff == 0:
				return resultSum
			case curDiff < 0:
				start++
			default:
				end--
			}
		}
	}
	return resultSum
}
