package array

import "sort"

// 散列表
func findRepeatNumber(nums []int) int {
	existNum := map[int]int{}
	for _, num := range nums {
		if _, ok := existNum[num]; !ok {
			existNum[num] = 0
		} else {
			return num
		}
	}
	return -1
}

// 排序
func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1
}

// 原地交换
func findRepeatNumber3(nums []int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == i {
			// 数字即在对应的索引位置
			i++
			continue
		}
		if nums[nums[i]] == nums[i] {
			// 重复
			return nums[i]
		}
		// 交换到正确的位置
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
