package main

//https://leetcode.com/problems/move-zeroes/description/

func moveZeroes(nums []int) {
	notNulCount :=0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[notNulCount] = nums[notNulCount], nums[i]
			notNulCount++
		}
	}
}

/*
func moveZeroes(nums []int) []int {
	var nuls []int
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nuls = append(nuls, 0)
		} else {
			i++
		}
	}
	nums = append(nums, nuls...)
	return nums
}
*/
/*
func moveZeroes(nums []int) {
	n := len(nums) - 1
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			j := i
			for nums[j] == 0 && j < n {
				j++
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
}
*/

/*
func moveZeroes(nums []int) {
	var nuls []int
	for i := 0; i < len(nums); {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nuls = append(nuls, 0)
		} else {
			i++
		}
	}
	nums = append(nums, nuls...)
}
*/