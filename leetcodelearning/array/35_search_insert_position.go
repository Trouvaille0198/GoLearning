package array

// Given a sorted array and a target value, return the index if the target is found.
// If not, return the index where it would be if it were inserted in order.
// You may assume no duplicates in the array.

// searchInsert 找到第一个比target大的 或者找到最后一个比target小的元素即可
func searchInsert(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if target < nums[mid] {
			if mid == 0 || target > nums[mid-1] {
				return mid
			}
			high = mid - 1
		} else if target > nums[mid] {
			if mid == len(nums)-1 || target < nums[mid+1] {
				return mid + 1
			}
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
