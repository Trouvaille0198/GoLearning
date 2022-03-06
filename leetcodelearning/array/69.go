package array

func mySqrt(x int) int {
	low, high := 0, x
	for low <= high {
		mid := low + (high-low)>>1
		if mid*mid < x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			low = mid + 1
		} else if mid*mid > x {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
