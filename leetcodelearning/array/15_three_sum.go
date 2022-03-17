package array

import "sort"

// Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.
// The solution set must not contain duplicate triplets.

// threeSum 排序+双指针
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < len(nums)-2; i++ {
		a := nums[i]
		if a > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			b, c := nums[l], nums[r]
			if a+b+c > 0 {
				r--
			} else if a+b+c < 0 {
				l++
			} else {
				res = append(res, []int{a, b, c})
				for l < r && nums[l] == b {
					l++
				}
				for r > l && nums[r] == c {
					r--
				}
			}
		}
	}
	return res
}
