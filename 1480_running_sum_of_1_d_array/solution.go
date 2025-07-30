package main

// https://leetcode.com/problems/running-sum-of-1d-array/description/

func runningSum(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	for i:=1; i< len(nums); i++{
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
