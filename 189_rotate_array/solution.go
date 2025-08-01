package main

func rotate(nums []int, k int) {
	br := len(nums) - (k % len(nums))
	copy(nums, append(nums[br:], nums[:br]...))
}
