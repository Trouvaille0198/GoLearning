package array

// Given a sorted array nums,
// remove the duplicates in-place such that each element appear only once and return the new length.
// Do not allocate extra space for another array,
// you must do this by modifying the input array in-place with O(1) extra memory.

// removeDuplicates 比较蠢的 一次遍历
func removeDuplicates(nums []int) int {
	dupLen := 0
	oriLen := len(nums)
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i+1], nums[i+2:]...) // 删除 i+1
			dupLen++
			// 此处不执行 i++
		} else {
			i++
		}
	}
	return oriLen - dupLen
}

// removeDuplicates2 把重复的元素移到最后面 （其实是覆盖掉）
func removeDuplicates2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	i, j := 0, 0
	for i < len(nums)-1 {
		for nums[j] == nums[i] {
			// 右移j至与i不同处
			j++
			if j == len(nums) {
				return i + 1
			}
		}
		nums[i+1] = nums[j]
		i++
	}
	return i + 1
}
