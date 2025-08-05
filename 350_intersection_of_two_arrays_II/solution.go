package main

import (
	"sort"
)

func intersect(nums1 []int, nums2 []int) []int {
	var res []int

	sort.Ints(nums1)
	sort.Ints(nums2)

	n1, n2 := len(nums1), len(nums2)
	i, j := 0, 0

	for i < n1 && j < n2 {
		if nums1[i] == nums2[j] {
			res = append(res, nums1[i])
			i++
			j++
		} else if nums1[i] > nums2[j]  {
			j++
		} else if nums2[j] > nums1[i]  {
			i++
		}
	}
	return res
}
