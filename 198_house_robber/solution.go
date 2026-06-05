package main

func rob(nums []int) int {
	n := len(nums)
	if n==1{
		return nums[0]
	}
	for i := 0; i < n; i++ {
		if i == 2 {
			nums[i] += nums[0]
		}
		if i > 2 {
			nums[i] += max(nums[i-2], nums[i-3])
		}
	}
	return max(nums[n-1], nums[n-2])
}
