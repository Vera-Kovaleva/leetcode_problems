package main

func missingNumber(nums []int) int {
	needSum, haveSum := 0, 0
	for i := 1; i <= len(nums); i++ {
		needSum += i
		haveSum += nums[i-1]
	}
	return needSum - haveSum
}
