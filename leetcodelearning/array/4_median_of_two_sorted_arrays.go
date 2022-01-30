package array

// There are two sorted arrays nums1 and nums2 of size m and n respectively.
// Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).
// You may assume nums1 and nums2 cannot be both empty.

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}

// findMedianSortedArrays 二分法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 假设 nums1 的长度小
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high := 0, len(nums1)
	k := (len(nums1) + len(nums2) + 1) >> 1 // 总长一半 向下取整
	// 接下来的操作 实际上就是找到第k小的元素 这道题的k值就是总长一半
	nums1Mid, nums2Mid := 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] {
			// nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] {
			// nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}

	midLeft, midRight := 0, 0
	switch {
	case nums1Mid == 0:
		midLeft = nums2[nums2Mid-1]
	case nums2Mid == 0:
		midLeft = nums1[nums1Mid-1]
	default:
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}

	if (len(nums1)+len(nums2))&1 == 1 {
		// 若为奇数 直接返回浮点形式
		return float64(midLeft)
	}
	// 若为偶数
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

// findMedianSortedArrays2 归并 将这两个有序数组合并后找中位数。
func findMedianSortedArrays2(nums1, nums2 []int) float64 {
	var result []int
	m, n := len(nums1), len(nums2)
	l1, l2 := 0, 0
	for l1 < m && l2 < n {
		if nums1[l1] < nums2[l2] {
			result = append(result, nums1[l1])
			l1++
		} else {
			result = append(result, nums2[l2])
			l2++
		}
	}
	result = append(result, nums1[l1:]...)
	result = append(result, nums2[l2:]...)
	length := m + n
	if length%2 == 1 {
		return float64(result[length/2])
	}
	mid1 := result[length/2]
	mid2 := result[length/2-1]
	return float64(mid1+mid2) / 2.0
}

// findMedianSortedArrays3 双指针 本质上还是归并
func findMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	length := m + n
	leftResult, result := 0, 0
	l1, l2 := 0, 0
	// 如果length为奇数，那问题转换为求第k小的数(k=length/2+1)
	// 所以循环k-1次，迭代right即可得到第k小的数
	// 如果length为偶数，那么问题转换为求第k-1小的数和第k小的数两个数的平均值
	for i := 0; i <= length/2; i++ {
		leftResult = result // 始终记录result左边的值
		if l1 < m && (l2 >= n || nums1[l1] < nums2[l2]) {
			result = nums1[l1]
			l1++
		} else {
			result = nums2[l2]
			l2++
		}
	}
	if length%2 == 1 {
		return float64(result)
	}
	return float64(leftResult+result) / 2.0
}
