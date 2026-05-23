package main

/*
https://leetcode.com/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150

Merge nums1 and nums2 into a single array sorted in non-decreasing order.

The final sorted array should not be returned by the function,
but instead be stored inside the array nums1. To accommodate this,
nums1 has a length of m + n, where the first m elements denote the
elements that should be merged, and the last n elements are set to 0
and should be ignored. nums2 has a length of n.
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := m-1, n-1
	for p := m + n - 1; p >= 0; p-- {
		if p2 < 0 {
			break
		}
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[p] = nums1[p1]
			p1--
		} else {
			nums1[p] = nums2[p2]
			p2--
		}
	}
}

/*
func merge(nums1 []int, m int, nums2 []int, n int) {
	p1, p2 := 0, 0
	nums1Copy := append([]int(nil), nums1[:m]...)
	for p := 0; p < m+n; p++ {
		if p1 >= m {
			copy(nums1[p:], nums2[p2:])
			return
		}
		if p2 >= n {
			copy(nums1[p:], nums1Copy[p1:])
			return
		}
		if nums1Copy[p1] < nums2[p2] {
			nums1[p] = nums1Copy[p1]
			p1++
		} else {
			nums1[p] = nums2[p2]
			p2++
		}
	}
	return
}
*/
