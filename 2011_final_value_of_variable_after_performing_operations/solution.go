package main

func finalValueAfterOperations(operations []string) int {
	x := 0

	for _, operation := range operations {
		if operation[1] == '+' {
			x++
		} else {
			x--
		}
	}

	return x
}
