package greedy

func canJump(nums []int) bool {
	maxIndex := 0
	for i := 0; i < len(nums); i++ {
		if i > maxIndex {
			return false
		}
		// 更新能跳到的最远位置
		if nums[i]+i > maxIndex {
			maxIndex = nums[i] + i
		}
	}
	return true
}

func canJump(nums []int) bool {
	dp := make([]bool, len(nums))
	dp[0] = true
	for i := 1; i < len(nums); i++ {
		flag := false
		for j := i - 1; j >= 0; j-- {
			if dp[j] && nums[j] >= i-j {
				flag = true
				break
			}
		}
		if flag {
			dp[i] = true
		}
	}
	return dp[len(nums)-1]
}
