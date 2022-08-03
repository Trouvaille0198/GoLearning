package array

func majorityElement(nums []int) int {
	sort.Ints(nums)
	midVal := nums[(0+len(nums)-1)>>1]
	return midVal
}

func majorityElement(nums []int) int {
	targetNum, count := nums[0], 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == targetNum {
			count++
		} else {
			count--
		}
		if count == -1 {
			targetNum, count = nums[i], 1
		}
	}
	return targetNum
}
