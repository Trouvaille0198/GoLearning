package array

func sortedSquares(nums []int) []int {
	ans := make([]int, len(nums))
	i := len(ans) - 1
	l, r := 0, len(nums)-1
	for l <= r {
		if nums[l]*nums[l] >= nums[r]*nums[r] {
			ans[i] = nums[l] * nums[l]
			l++
		} else {
			ans[i] = nums[r] * nums[r]
			r--
		}
		i--
	}
	return ans
}
