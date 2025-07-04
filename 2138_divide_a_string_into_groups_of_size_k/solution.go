package main

import "strings"

func divideString(s string, k int, fill byte) []string {
	var result []string

	start := 0
	n := len(s)

	lastPiece := ""

	for start < n {
		end := start + k
		if end > n {
			end = n
			lastPiece = s[start:end]
			start = end
		} else {
			result = append(result, s[start:end])
			start = end
		}
	}

	if len(lastPiece) > 0 {
		lastPiece += strings.Repeat(string(fill), k-len(lastPiece))
		result = append(result, lastPiece)
	}

	return result
}
