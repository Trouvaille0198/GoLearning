
func findKthLargest(nums []int, k int) int {
	var partition func(nums []int, l, r int) int
	partition = func(nums []int, l, r int) int {
		// pivot := nums[r]
		// for i := l; i < r; i++ {
		// 	if nums[i] > pivot {
		// 		nums[l], nums[i] = nums[i], nums[r]
		// 		l++
		// 	}
		// }
		// nums[l], nums[r] = nums[r], nums[l]
		// return l
		pivotPos := true // 基准值位置 true为前
		i, j := l, r
		for i < j {
			if nums[i] < nums[j] {
				// 逆序
				pivotPos = !pivotPos
				nums[i], nums[j] = nums[j], nums[i]
			}
			if pivotPos {
				j--
			} else {
				i++
			}
		}
		return i
	}

	left := 0
	right := len(nums) - 1

	for left < right {
		p := partition(nums, left, right)
		if p+1 == k {
			return nums[p]
		} else if p+1 < k {
			left = p + 1
		} else {
			right = p - 1
		}
	}
	return nums[right]
}
