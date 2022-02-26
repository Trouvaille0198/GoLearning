package array

import "sort"

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

// 快排思想
func getLeastNumbers2(arr []int, k int) []int {
	var quickSort func(nums []int, l, r, x int)
	quickSort = func(nums []int, l, r, x int) {
		if l > r {
			return
		}
		i, j := l, r
		pivotPos := true
		for i < j {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
				pivotPos = !pivotPos
			}
			if pivotPos {
				j--
			} else {
				i++
			}
		}
		num := i - l + 1 // 本次划分 可以保证划分正确的个数为num
		if x == num {
			// 正好
			return
		} else if x < num {
			quickSort(nums, l, i-1, x)
		} else {
			quickSort(nums, i+1, r, x-num)
		}
	}
	if k == 0 {
		return []int{}
	}
	quickSort(arr, 0, len(arr)-1, k)
	return arr[:k]
}

// 堆排
func getLeastNumbers3(arr []int, k int) []int {
	// 大顶堆
	var heapify func(nums []int, root, end int)
	heapify = func(nums []int, root, end int) {
		// 大顶堆堆化，堆顶值小一直下沉
		for {
			// 左孩子节点索引
			child := root*2 + 1
			// 越界跳出
			if child > end {
				return
			}
			// 比较左右孩子，取大值，否则child不用++
			if child < end && nums[child] <= nums[child+1] {
				child++
			}
			// 如果父节点已经大于左右孩子大值，已堆化
			if nums[root] > nums[child] {
				return
			}
			// 孩子节点大值上冒
			nums[root], nums[child] = nums[child], nums[root]
			// 更新父节点到子节点，继续往下比较，不断下沉
			root = child
		}
	}
	end := len(arr) - 1
	// 从最后一个非叶子节点开始堆化
	for i := end / 2; i >= 0; i-- {
		heapify(arr, i, end)
	}
	// 依次弹出元素，然后再堆化，相当于依次把最大值放入尾部
	for i := end; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		end--
		heapify(arr, 0, end)
	}
	return arr[:k]
}
