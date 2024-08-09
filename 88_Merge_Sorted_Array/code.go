package mergesortedarray

import "slices"

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < len(nums2); i++ {
		nums1[m] = nums2[i]
		m++
	}
	slices.Sort(nums1)
}
