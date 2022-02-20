package array

// 前缀和
func subarraySum(nums []int, k int) int {
	count := 0
	curSum := 0 // 记录每次循环当前的前缀和
	hash := map[int]int{
		0: 1, // 前缀和为0 出现过1次了
	}
	for i := 0; i < len(nums); i++ {
		curSum += nums[i]
		if hash[curSum-k] > 0 {
			// 当前前缀和减去k的结果也在之前的前缀和表中 代表找到一个和为k的连续子数组
			count += hash[curSum-k]
		}
		hash[curSum]++ // 记录一个前缀和
	}
	return count
}
