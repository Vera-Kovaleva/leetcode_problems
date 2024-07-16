package main

func climbStairs(n int) int {
	ways := make(map[int]int)
	ways[0] = 1
	ways[1] = 1
	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]
	}
	return ways[n]
}
