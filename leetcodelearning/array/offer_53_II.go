package array

// 遍历 出现前后差为2则返回
func missingNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i+1]-nums[i] > 1 {
			return i + 1
		}
	}
	if nums[len(nums)-1] == len(nums)-1 {
		return len(nums)
	} else {
		return 0
	}
}

// 等差数列求和
func missingNumber2(nums []int) int {
	n := len(nums)
	totalSum := n * (n + 1) / 2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return totalSum - sum
}

// 二分
func missingNumber3(nums []int) int {
	low, high := 0, len(nums)-1
	for low <= high {
		mid := (low + high) >> 1
		if mid == nums[mid] {
			// 缺失值在右边
			if mid == len(nums)-1 || nums[mid+1]-nums[mid] > 1 {
				return nums[mid] + 1
			}
			low = mid + 1
		} else {
			// 缺失值在左边
			if mid == 0 || nums[mid]-nums[mid-1] > 1 {
				return nums[mid] - 1
			}
			high = mid - 1
		}
	}
	return -1
}
