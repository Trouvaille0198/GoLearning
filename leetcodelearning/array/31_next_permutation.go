package array

// Implement next permutation,
// which rearranges numbers into the lexicographically next greater permutation of numbers.
// If such an arrangement is not possible, it must rearrange it as the lowest possible order
// (i.e., sorted in ascending order).
// The replacement must be in place and use only constant extra memory.

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	smallIndex, largeIndex := 0, 0
	// 寻找较小数和较大数
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < nums[i+1] {
			if i+1 == len(nums)-1 {
				// 若i+1后没有区间 直接交换
				nums[i], nums[i+1] = nums[i+1], nums[i]
				return
			}
			isReversed := true
			for j := i + 1; j < len(nums)-1; j++ {
				// 判断i+1后是否降序的
				if nums[j] < nums[j+1] {
					isReversed = false
					break
				}
			}
			if isReversed {
				smallIndex = i // 确认较小数
				// 找较大数
				for j := len(nums) - 1; j >= i+1; j-- {
					if nums[j] > nums[i] {
						largeIndex = j // 确认较大数
						break
					}
				}
			}
		}
	}
	nums[smallIndex], nums[largeIndex] = nums[largeIndex], nums[smallIndex]

	start, end := 0, len(nums)-1 // 默认值 若找不到较大数和较小数 代表数组没有下一个排列 就将整个数组全部升序
	if smallIndex != largeIndex {
		// 将i+1后的区间调为正序
		start, end = smallIndex+1, len(nums)-1
	}

	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
