package array

import (
	"math"
	"sort"
)

// 很烂的方法
func isStraight(nums []int) bool {
	sort.Ints(nums)
	zeroCount := 0
	for _, num := range nums {
		if num == 0 {
			zeroCount++
		}
	}
	one, two := 0, 0
	for i := zeroCount + 1; i < 5; i++ {
		{
			diff := nums[i] - nums[i-1]
			if diff == 1 {
				continue
			} else if diff == 2 {
				one++
			} else if diff == 3 {
				two++
			} else {
				return false
			}
		}
	}
	if (one == 0 && two == 0) || (one <= zeroCount && two == 0) || (one == 0 && two*2 <= zeroCount) {
		return true
	} else {
		return false
	}
}

// 哈希+遍历
func isStraight2(nums []int) bool {
	hash := make(map[int]struct{})
	minVal, maxVal := math.MaxInt, math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			// 把0跳过
			if _, ok := hash[nums[i]]; ok {
				// 有重复 必不可能构成顺子
				return false
			}
			hash[nums[i]] = struct{}{}
			if nums[i] < minVal {
				minVal = nums[i]
			}
			if nums[i] > maxVal {
				maxVal = nums[i]
			}
		}
	}
	if maxVal-minVal < 5 {
		return true
	}
	return false
}
