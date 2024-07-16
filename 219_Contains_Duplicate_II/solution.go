package main

func containsNearbyDuplicate(nums []int, k int) bool {
	numsAndIndexes := make(map[int][]int)

	for i, num := range nums {
		prevIndexes := numsAndIndexes[num]
		if prevIndexes == nil {
			prevIndexes = []int{i}
		} else {
			prevIndexes = append(prevIndexes, i)
		}
		numsAndIndexes[num] = prevIndexes
	}

	for _, indexes := range numsAndIndexes {
		for i := 1; i < len(indexes); i++ {
			if indexes[i]-indexes[i-1] <= k {
				return true
			}
		}
	}

	return false
}
