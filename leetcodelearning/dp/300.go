package dp

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		// nums[i] 最大
		if nums[i] > dp[len(dp)-1] {
			dp = append(dp, nums[i])
			continue
		}
		low, high := 0, len(dp)-1
		for low <= high {
			mid := low + (high-low)>>1
			if nums[i] <= dp[mid] {
				if mid == 0 || nums[i] > dp[mid-1] {
					// 找到第一个大于等于nums[i]的元素
					dp[mid] = nums[i] // 替换
					break
				}
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
	return len(dp)
}
