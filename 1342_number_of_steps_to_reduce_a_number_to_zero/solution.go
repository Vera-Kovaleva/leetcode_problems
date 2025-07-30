package main

func numberOfSteps(num int) int {
	numOfSteps := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
			numOfSteps++
		} else if num > 0 {
			num--
			numOfSteps++
		}
	}
	return numOfSteps
}
