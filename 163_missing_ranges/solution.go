package main

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	n := len(nums)
	res := make([][]int, 0)

	if n == 0 {
		res = append(res, []int{lower, upper})
		return res
	}

	if lower < nums[0] {
		res = append(res, []int{lower, nums[0] - 1})
	}

	for i := 0; i < n-1; i++ {
		if !(nums[i+1]-nums[i] <= 1) {
			res = append(res, []int{nums[i] + 1, nums[i+1] - 1})
		}
	}

	if upper > nums[n-1] {
		res = append(res, []int{nums[n-1] + 1, upper})
	}

	return res
}
