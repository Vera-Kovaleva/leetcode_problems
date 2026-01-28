package main

func convertToTitle(columnNumber int) string {
	res := ""
	alphabet := "ABCDEFGHIJKLMNOPQRESTUVWYZ"

	for columnNumber > 0 {
		columnNumber--
		res = string(alphabet[columnNumber%26]) + res
		columnNumber /= 26
	}

	return res
}
