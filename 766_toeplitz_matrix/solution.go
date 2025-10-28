package main

/*
func isToeplitzMatrix(matrix [][]int) bool {
	diagonals := make(map[int]int)
	for i, row := range matrix {
		for j, element := range row {
			if diagValue, ok := diagonals[i-j]; !ok {
				diagonals[i-j]=element
			} else if diagValue!=element {
				return false
			}
		}
	}
    return true
}
*/

func isToeplitzMatrix(matrix [][]int) bool {
	for r, _ := range matrix {
		for c, _ := range matrix[r] {
			if r > 0 && c > 0 && matrix[r-1][c-1] != matrix[r][c] {
				return false
			}
		}
	}
	return true
}
