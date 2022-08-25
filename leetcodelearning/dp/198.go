package dp

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	if len(nums) == 3 {
		return max(nums[1], nums[2]+nums[0])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[2] + dp[0]
	res := max(dp[1], dp[2])
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-2], dp[i-3]) + nums[i]
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res

}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	var fir, sec, cur int
	fir, sec = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		cur = max(sec, fir+nums[i])
		fir = sec
		sec = cur
	}

	return cur

}
