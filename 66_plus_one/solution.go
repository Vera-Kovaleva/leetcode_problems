package main

func plusOne(digits []int) []int {
	n := len(digits) - 1
	digits[n]++
	for i := n; i > 0; i-- {
		if digits[i] != 10 {
			return digits
		}
		digits[i] = 0
		digits[i-1]++
	}
	if digits[0] == 10 {
		digits = append([]int{1, 0}, digits[1:]...)
	}
	return digits
}
