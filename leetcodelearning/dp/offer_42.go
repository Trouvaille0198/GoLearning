package dp

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxResult := dp[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if dp[i] > maxResult {
			maxResult = dp[i]
		}
	}
	return maxResult
}

func maxSubArray2(nums []int) int {
	maxResult := nums[0]
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if pre > 0 {
			pre = pre + nums[i]
		} else {
			pre = nums[i]
		}
		if pre > maxResult {
			maxResult = pre
		}
	}
	return maxResult
}
