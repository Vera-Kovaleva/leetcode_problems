package main

import "math"

/*
func getSneakyNumbers(nums []int) []int {
	ans := []int{}
	counter := make(map[int]int)
	for _, num := range nums {
		counter[num]++
		if counter[num] >= 2 {
			ans = append(ans, num)
		}
		if len(ans) == 2 {
			return ans
		}
	}
	return ans
}
*/

/*
func getSneakyNumbers(nums []int) []int {
	n := len(nums) - 2
	y := 0

	for _, num := range nums {
		y ^= num
	}

	for i := 0; i < n; i++ {
		y ^= i
	}

	lowBit := y & -y

	x1, x2 := 0, 0

	for _, x := range nums {
		if x&lowBit != 0 {
			x1 ^= x
		} else {
			x2 ^= x
		}
	}

	for i := 0; i < n; i++ {
		if i&lowBit != 0 {
			x1 ^= i
		} else {
			x2 ^= i
		}
	}

	return []int{x1, x2}
}
*/

func getSneakyNumbers(nums []int) []int {
	/*
		{ x1 + x2 = sumDiff = realSum - expectedSum
		{ x1^2 + x2^2 = sumSquardedDiff = realSquaredSum - expectedSquaredSum
	*/

	n := len(nums) - 2

	expectedSum := (n) * (n - 1) / 2
	expectedSquaredSum := n * (n - 1) * (2*n - 1) / 6

	sum, squaredSum := 0, 0
	for _, x := range nums {
		sum += x
		squaredSum += x * x
	}

	sumDiff := sum - expectedSum
	sumSqDiff := squaredSum - expectedSquaredSum

	discriminant := 2*sumSqDiff - sumDiff*sumDiff
	root := int(math.Sqrt(float64(discriminant)))

	x1 := (sumDiff - root) / 2
	x2 := sumDiff - x1

	return []int{x1, x2}

}
