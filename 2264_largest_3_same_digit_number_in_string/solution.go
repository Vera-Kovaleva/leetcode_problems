package main

func largestGoodInteger(num string) string {
	best := ""
	for i := 0; i+2 < len(num); i++ {
		if num[i] == num[i+1] && num[i+1] == num[i+2] {
			if num[i:i+3] > best {
				best = num[i : i+3]
			}
		}
	}

	return best
}
