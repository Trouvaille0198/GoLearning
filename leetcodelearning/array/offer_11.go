package array

// O(n) 遍历
func minArray(numbers []int) int {
	if len(numbers) == 1 {
		return numbers[0]
	}
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			return numbers[i+1]
		}
	}
	return numbers[0]
}

// 二分
func minArray2(numbers []int) int {
	low := 0
	high := len(numbers) - 1
	for low < high {
		pivot := low + (high-low)/2
		if numbers[pivot] < numbers[high] {
			high = pivot
		} else if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return numbers[low]
}
