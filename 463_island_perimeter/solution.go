package main

func islandPerimeter(grid [][]int) int {
	res := 0
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				cur := 4
				if i > 0 {
					cur -= grid[i-1][j]
				}
				if i < rows-1 {
					cur -= grid[i+1][j]
				}
				if j > 0 {
					cur -= grid[i][j-1]
				}
				if j < cols-1 {
					cur -= grid[i][j+1]
				}
				res += cur
			}
		}
	}
	return res
}
