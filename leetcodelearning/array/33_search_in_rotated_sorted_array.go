package array

// Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.
// (i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).
// You are given a target value to search. If found in the array return its index, otherwise return -1.
// You may assume no duplicate exists in the array.
// Your algorithm’s runtime complexity must be in the order of O(log n).

// 二分法
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	low, high := 0, len(nums)-1
	for low <= high {
		// 开始二分
		mid := low + (high-low)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > nums[high] {
			// 中间值大于右边值 左边必然有序
			if nums[low] <= target && target < nums[mid] {
				// 落在左边部分
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			// 中间值小于右边值 右边必然有序
			if nums[mid] < target && target <= nums[high] {
				// 落在右边部分
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
