package main

import (
	"strconv"
)

//https://leetcode.com/problems/base-7/

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	negative := false
	if num < 0 {
		negative = true
		num = -num
	}
	var result string

	for num > 0 {
		ost := num % 7
		result = strconv.Itoa(ost) + result
		num = num / 7
	}

	if negative {
		result = "-" + result
	}

	return result
}
