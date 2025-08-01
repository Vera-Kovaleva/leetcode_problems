package main

func containsDuplicate(nums []int) bool {
	controller := make(map[int]struct{})
	for _, num := range nums {
		_, exists := controller[num]
		if exists {
			return true
		}
		controller[num] = struct{}{}
	}
	return false
}