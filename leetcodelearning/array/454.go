package array

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	hash := make(map[int]int)
	count := 0
	for _, a := range nums1 {
		for _, b := range nums2 {
			hash[a+b]++
		}
	}
	for _, c := range nums3 {
		for _, d := range nums4 {
			if _, ok := hash[-(c + d)]; ok {
				count += hash[-(c + d)]
			}
		}
	}
	return count
}
