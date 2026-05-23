package main

func binaryGap(n int) int {
	last := -1
	maxGap := 0
	pos := 0

	for n > 0 {
		if n&1 == 1 {
			if last != -1 {
				if pos-last > maxGap {
					maxGap = pos - last
				}
			}
			last = pos
		}
		n >>= 1
		pos++
	}

	return maxGap
}
