package array

import "sort"

// Given an array nums of n integers and an integer target
// are there elements a, b, c, and d in nums such that a + b + c + d = target?
// Find all unique quadruplets in the array which gives the sum of target.
// The solution set must not contain duplicate quadruplets.

// fourSum 排序+双指针 外加一层循环
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	curNums := make([][]int, 0)

	for j := 0; j < len(nums); j++ {
		if j > 0 && nums[j] == nums[j-1] {
			continue
		}
		for i := j + 2; i < len(nums); i++ {
			start, end := j+1, len(nums)-1
			// 考虑i跟前一个i相同的情况
			if i > j+2 && nums[i] == nums[i-1] {
				start = i - 1
			}

			for start < i && end > i {
				// 去重
				if start > j+1 && nums[start] == nums[start-1] {
					start++
					continue
				}
				if end < len(nums)-1 && nums[end] == nums[end+1] {
					end--
					continue
				}

				curSum := nums[j] + nums[start] + nums[i] + nums[end]
				switch {
				case curSum == target:
					curNums = append(curNums, []int{nums[j], nums[start], nums[i], nums[end]})
					start++
					end--
				case curSum < target:
					start++
				default:
					end--
				}
			}
		}
	}
	return curNums
}
