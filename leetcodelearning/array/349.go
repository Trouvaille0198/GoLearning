package array

func intersection(nums1 []int, nums2 []int) (res []int) {
	hash := make(map[int]struct{}, 0)
	for i := 0; i < len(nums1); i++ {
		hash[nums1[i]] = struct{}{}
	}
	for i := 0; i < len(nums2); i++ {
		if _, ok := hash[nums2[i]]; ok {
			res = append(res, nums2[i])
			delete(hash, nums2[i]) // 去重
		}
	}
	return
}
