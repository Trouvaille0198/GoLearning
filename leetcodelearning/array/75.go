package array

// 单指针
func sortColors(nums []int) {
	k := 0
	// 把所有0放到开头
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
	// 把所有1放到0后
	for i := k; i < len(nums); i++ {
		if nums[i] == 1 {
			nums[i], nums[k] = nums[k], nums[i]
			k++
		}
	}
}

// 双指针
func sortColors(nums []int) {
	p0, p1 := 0, 0
	for i, num := range nums {
		if num == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			if p0 < p1 {
				nums[i], nums[p1] = nums[p1], nums[i]
			}
			p0++
			p1++
		} else if num == 1 {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
		}
	}
}

// 双指针2
func sortColors(nums []int) {
	p0, p2 := 0, len(nums)-1
	for i := 0; i <= p2; i++ {
		for i <= p2 && nums[i] == 2 {
			nums[i], nums[p2] = nums[p2], nums[i]
			p2--
		}
		if nums[i] == 0 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
