package array

// 贪心 + 前缀和
func maxNonOverlapping(nums []int, target int) int {
	count, curSum := 0, 0
	hash := map[int]int{0: 1}
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		if hash[curSum-target] > 0 {
			count++ // 只记录一个
			// 清空前缀和记录
			hash = map[int]int{}
			curSum = 0
		}
		hash[curSum]++
	}
	return count
}
