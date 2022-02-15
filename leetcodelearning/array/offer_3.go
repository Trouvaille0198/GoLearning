package array

import "sort"

func findRepeatNumber(nums []int) int {
	existNum := map[int]int{}
	for _, num := range nums {
		if _, ok := existNum[num]; !ok {
			existNum[num] = 0
		} else {
			return num
		}
	}
	return -1
}

func findRepeatNumber2(nums []int) int {
	sort.Ints(nums)
}
