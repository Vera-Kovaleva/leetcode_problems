package main

func countValidSelections(nums []int) int {
	count := 0
	sum, leftSum := 0, 0
	for _, num := range nums {
		sum += num
	}

	for _, num := range nums {
		leftSum += num
		if num == 0 {
			diff := 2*leftSum - sum
			if diff == 1 || diff == -1 {
				count += 1
			}
			if diff == 0 {
				count += 2
			}
		}
	}
	return count
}
