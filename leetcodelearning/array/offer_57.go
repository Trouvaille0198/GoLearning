package array

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]+nums[j] < target {
			i++
		} else if nums[i]+nums[j] > target {
			j--
		} else {
			return []int{nums[i], nums[j]}
		}
	}
	return []int{}
}
