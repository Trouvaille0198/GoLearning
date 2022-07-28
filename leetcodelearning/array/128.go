package array

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numSet := make(map[int]bool, 0)
	for _, num := range nums {
		numSet[num] = true
	}
	maxLen := 1
	for num := range numSet {
		curNum := num
		curLen := 1
		if !numSet[num-1] {
			for numSet[curNum+1] {
				curNum++
				curLen++
			}
			if curLen > maxLen {
				maxLen = curLen
			}
		}
	}
	return maxLen
}
