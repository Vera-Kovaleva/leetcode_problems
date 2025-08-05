package main

func twoSum(nums []int, target int) []int {
	n := len(nums)
	check := make(map[int]int)
	for i := 0; i < n; i++ {
		cur := nums[i]
		value, ok := check[target-cur]
		if ok {
			return []int{value, i}
		}
		check[cur] = i
	}
	return []int{}
}
