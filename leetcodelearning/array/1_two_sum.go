package array

// Given an array of integers, return indices of the two numbers such that they add up to a specific target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // 键为数字 值为下标
	for i := 0; i < len(nums); i++ {
		if index, ok := numMap[target-nums[i]]; ok {
			return []int{i, index}
		}
		numMap[nums[i]] = i
	}
	return nil
}
