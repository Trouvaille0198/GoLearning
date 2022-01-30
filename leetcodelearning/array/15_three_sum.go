package array

import "sort"

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
// The solution set must not contain duplicate triplets.

// threeSum 排序+双指针
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0)

	for index := 1; index < len(nums)-1; index++ {
		start, end := 0, len(nums)-1
		if index > 1 && nums[index] == nums[index-1] {
			// 考虑index重复的情况
			start = index - 1
		}

		for start < index && end > index {
			// 跳过重复项
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}

			addNum := nums[start] + nums[end] + nums[index]
			switch {
			case addNum == 0:
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			case addNum > 0:
				end--
			default:
				start++
			}
		}
	}
	return result
}
