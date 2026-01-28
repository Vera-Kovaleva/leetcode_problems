package main

func minCost(colors string, neededTime []int) int {
	minTime := 0
	n := len(neededTime)
	for i := 0; i < n-1; i++ {
		curMaxi, sum := 0, 0
		for i < n-1 && colors[i] == colors[i+1] {
			sum += neededTime[i]
			if neededTime[i] > curMaxi {
				curMaxi = neededTime[i]
			}
			i++
		}
		sum += neededTime[i]
		if neededTime[i] > curMaxi {
			curMaxi = neededTime[i]
		}
		minTime += sum - curMaxi
	}
	return minTime
}
