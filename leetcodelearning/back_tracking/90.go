package back_tracking

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	subset := make([]int, 0)
	var dfs func(startIndex int)
	dfs = func(startIndex int) {
		res = append(res, append([]int{}, subset...))
		for i := startIndex; i < len(nums); i++ {
			if i > startIndex && nums[i] == nums[i-1] {
				// 去重
				continue
			}
			subset = append(subset, nums[i])
			dfs(i + 1)
			subset = subset[:len(subset)-1]
		}
	}
	dfs(0)
	return res
}
