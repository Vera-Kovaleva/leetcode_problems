package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	poisoned_seconds := 0

	for i := 0; i < len(timeSeries)-1; i++ {
		poisoned_seconds += min(timeSeries[i]+duration, timeSeries[i+1]) - timeSeries[i]
	}
	poisoned_seconds += duration

	return poisoned_seconds
}