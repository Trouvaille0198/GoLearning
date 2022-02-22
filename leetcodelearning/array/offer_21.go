package array

// 快慢指针
func exchange(nums []int) []int {
	// i始终在第一个偶数上 j往前探索找奇数
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast]%2 == 1 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
		fast++
	}
	return nums
}

// 比较难看的快慢指针
func exchange2(nums []int) []int {
	// i始终在第一个偶数上 j往前探索找奇数
	i, j := 0, 0
	for i < len(nums) {
		if nums[i]%2 == 0 {
			// 若是偶数
			for j < len(nums) && nums[j]%2 == 0 {
				// 找到下一个奇数
				j++
			}
			if j == len(nums) {
				// 后面没有奇数了 直接返回
				return nums
			}
			// 交换
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			// 若是偶数
			j++ // j跟着i往前走一格
		}
		i++
	}
	return nums
}

// 头尾指针
func exchange3(nums []int) []int {
	low, high := 0, len(nums)-1

	for low <= high {
		for low < len(nums) && nums[low]%2 == 1 {
			// low往前走 直到遇到偶数
			low++
		}
		for high >= 0 && nums[high]%2 == 0 {
			// high往后退 直到遇到奇数
			high--
		}
		if low < high {
			nums[low], nums[high] = nums[high], nums[low]
		}
	}
	return nums
}
