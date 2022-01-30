package array

// Given an array of integers nums sorted in ascending order,
// find the starting and ending position of a given target value.
// Your algorithm’s runtime complexity must be in the order of O(log n).
// If the target is not found in the array, return [-1, -1].

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	firstIndex, lastIndex := -1, -1
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if target > nums[mid] {
			low = mid + 1
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			firstIndex, lastIndex = mid, mid
			for firstIndex >= low && nums[firstIndex] == target {
				firstIndex--
			}
			for lastIndex <= high && nums[lastIndex] == target {
				lastIndex++
			}
			return []int{firstIndex + 1, lastIndex - 1}

		}
	}
	return []int{-1, -1}
}
