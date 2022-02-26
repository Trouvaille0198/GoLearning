package array

import (
	"sort"
	"strconv"
)

// 获取两个数字拼接后的结果
func plusNum(a, b int) (res int) {
	res, _ = strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return
}

func compare(a, b int) bool {
	return plusNum(a, b) < plusNum(b, a)
}

func minNumber(nums []int) string {
	var res string
	// 选择排序
	sort.Slice(nums, func(i, j int) bool {
		return plusNum(nums[i], nums[j]) < plusNum(nums[j], nums[i])
	})
	for _, item := range nums {
		res += strconv.Itoa(item)
	}
	return res
}

// 自己写个快排
func minNumber2(nums []int) string {
	var quickSort func(nums []int, l, r int)
	quickSort = func(nums []int, l, r int) {
		pivotPos := true // 基准值位置 true为前
		i, j := l, r
		for i < j {
			if !compare(nums[i], nums[j]) {
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
		if l < i-1 {
			quickSort(nums, l, i-1)
		}
		if r > i+1 {
			quickSort(nums, i+1, r)
		}
	}
	quickSort(nums, 0, len(nums)-1)
	var res string
	for _, item := range nums {
		res += strconv.Itoa(item)
	}
	return res
}
