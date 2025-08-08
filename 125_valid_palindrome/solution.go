package main

import (
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	leftIndex, rightIndex := 0, len(s)-1
	for leftIndex < rightIndex {
		leftSymb := s[leftIndex]
		rightSymb := s[rightIndex]
		if (leftSymb >= 'a' && leftSymb <= 'z' || leftSymb >= '0' && leftSymb <= '9') && (rightSymb >= 'a' && rightSymb <= 'z' || rightSymb >= '0' && rightSymb <= '9') {
			if leftSymb != rightSymb {
				return false
			}
			leftIndex++
			rightIndex--
		}
		if !(leftSymb >= 'a' && leftSymb <= 'z' || leftSymb >= '0' && leftSymb <= '9') {
			leftIndex++
		}
		if !(rightSymb >= 'a' && rightSymb <= 'z' || rightSymb >= '0' && rightSymb <= '9') {
			rightIndex--
		}
	}
	return true
}
