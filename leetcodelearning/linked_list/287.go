package linked_list

func findDuplicate(nums []int) int {
	slow, fast := 0, 0
	for fast != slow || fast == 0 {
		// slow走一步 fast走两步 直到相遇
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	start := 0 // 设start从头走
	for start != slow {
		start = nums[start]
		slow = nums[slow]
	}
	return slow
}
