package main

func isTheSame(first [][]int, second [][]int) bool {
	if len(first) != len(second) {
		return false
	}
	n := len(first)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if first[i][j] != second[i][j] {
				return false
			}
		}
	}
	return true
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		for j, k := 0, n-1; j < k; j, k = j+1, k-1 {
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}

func findRotation(mat [][]int, target [][]int) bool {
	for i:=0; i<4;i++ {
		rotate(mat)
		if isTheSame(mat, target) {
			return true
		}
	}
	return false
}
