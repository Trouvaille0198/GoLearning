package array

// Given an array nums and a value val, remove all instances of that value in-place and return the new length.
// Do not allocate extra space for another array,
// you must do this by modifying the input array in-place with O(1) extra memory.

// The order of elements can be changed. It doesn’t matter what you leave beyond the new length.

// removeDuplicates 将重复元素与当前指针交换
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return j
}
