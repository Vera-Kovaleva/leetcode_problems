package main

import "math/bits"

func countPrimeSetBits(left int, right int) int {
	isPrime := [21]bool{false, false, true, true, false, true, false, true, false, false, false, true, false, true, false, false, false, true, false, true, false}

	result := 0
	for n := left; n <= right; n++ {
		if isPrime[bits.OnesCount(uint(n))] {
			result++
		}
	}
	return result
}
