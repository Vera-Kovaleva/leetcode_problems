package main

func isValidSudoku(board [][]byte) bool {
	var rows, columns, squares [9][9]bool
	for i, curRow := range board {
		for j, curValue := range curRow {
			if curValue != '.' {
				k := int(curValue) - 49
				if rows[i][k] || columns[j][k] || squares[i/3*3+j/3][k] {
					return false
				}
				rows[i][k], columns[j][k], squares[i/3*3+j/3][k] = true, true, true
			}
		}
	}
	return true

}
