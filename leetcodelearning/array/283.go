package array

func moveZeroes(nums []int) {
	var i, j = 0, 0
	for j < len(nums) {
		for nums[j] == 0 {
			// j移动到第一个非零位置
			j++
			if j == len(nums) {
				return
			}
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j++
	}
}
