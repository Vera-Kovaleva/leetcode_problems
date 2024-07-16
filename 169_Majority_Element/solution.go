package main

	
import (
    "fmt"
    "slices"
)

/*
func majorityElement(nums []int) int {
    howManyNumbers := make(map[int]int)
	var n int
	n = len(nums)
	for _, num := range nums {
		howManyNumbers[num+10000000000] += 1
		if (howManyNumbers[num+10000000000]>n/2) {
			return num;
		}
	}
	return 0;
}
*/

func majorityElement(nums []int) int {
	n := len(nums)
	slices.Sort(nums)
	return nums[n/2];
}