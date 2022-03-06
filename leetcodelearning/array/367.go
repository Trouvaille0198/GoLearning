package array

func isPerfectSquare(num int) bool {
	low, high := 0, num
	for low <= high {
		mid := low + (high-low)>>1
		if mid*mid < num {
			low = mid + 1
		} else if mid*mid > num {
			high = mid - 1
		} else {
			return true
		}
	}
	return false
}
