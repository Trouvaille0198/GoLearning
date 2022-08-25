package array

func productExceptSelf(nums []int) []int {
	l, r, res := make([]int, len(nums)), make([]int, len(nums)), make([]int, len(nums))
	// l[i]为i左边元素的乘积
	// r[i]为i右边元素的乘积
	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	r[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = l[i] * r[i]
	}
	return res
}

func productExceptSelf2(nums []int) []int {
	res := make([]int, len(nums))
	// res[i]为i左边元素的乘积
	res[0] = 1
	for i := 1; i < len(nums); i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	r := 1 // i右侧元素乘积
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = res[i] * r
		r *= nums[i]
	}
	return res
}
