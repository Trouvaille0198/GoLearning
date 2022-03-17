package back_tracking

func permute(nums []int) [][]int {
	res := make([][]int, 0)
	var bt func(nums, path []int)
	bt = func(nums, path []int) {
		if len(nums) == 0 {
			p := make([]int, len(path))
			copy(p, path) // 拷贝
			res = append(res, p)
			return
		}
		n := len(nums)
		for i := 0; i < n; i++ {
			cur := nums[i]
			path = append(path, cur)
			nums = append(nums[:i], nums[i+1:]...)
			bt(nums, path)
			// 还原
			nums = append(nums[:i], append([]int{cur}, nums[i:]...)...)
			path = path[:len(path)-1]
		}
	}
	bt(nums, []int{})
	return res
}
