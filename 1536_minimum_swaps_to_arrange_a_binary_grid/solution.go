package main

func minSwaps(grid [][]int) int {
	swaps := 0

	n := len(grid)
	positions := make([]int, n)

	for i := 0; i < n; i++ {
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 1 {
				positions[i] = j
				break
			}
		}
	}

	for i := 0; i < n; i++ {
		k := -1
		for j := i; j < n; j++ {
			if positions[j] <= i {
				swaps += j - i
				k = j
				break
			}
		}

		if k != -1 {
			for j := k; j > i; j-- {
				positions[j], positions[j-1] = positions[j-1], positions[j]
			}
		} else {
			return -1
		}
	}

	return swaps
}
