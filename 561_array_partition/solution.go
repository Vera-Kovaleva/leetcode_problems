package main

import "slices"

//https://leetcode.com/problems/array-partition/description/

func arrayPairSum(nums []int) int {
	slices.Sort(nums)
	res := 0
	for i := 0; i < len(nums); i += 2 {
		res += nums[i]
	}
	return res
}
