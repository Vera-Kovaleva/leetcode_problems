package main

import "math"

func myAtoi(s string) int {
	var res int

	lenS := len(s)
	idx := 0

	for idx < lenS && (s[idx] == ' ' || s[idx] == '\t') {
		idx++
	}
	if idx == lenS {
		return 0
	}

	sign := 1
	if s[idx] == '+' {
		sign = 1
		idx++
	} else if s[idx] == '-' {
		sign = -1
		idx++
	}

	for idx < lenS {
		symb := s[idx]

		if '0' <= symb && symb <= '9' {
			res = res*10 + int(symb-'0')
			if sign * res > math.MaxInt32 {
				return math.MaxInt32
			}
	
			if sign * res < math.MinInt32 {
				return math.MinInt32
			}
			idx++
		} else {
			return res * sign
		}
	}

	return res * sign
}

/*
func myAtoi2(s string) int {
	var res int

	lenS := len(s)
	idx := 0

	symb := s[idx]
	for idx < lenS && (symb == ' ' || symb == '\t') {
		idx++
	}

	if idx == lenS {
		return 0
	}

	sign := 1
	if s[idx] == '+' {
		sign = 1
		idx++
	} else if s[idx] == '-' {
		sign = -1
		idx++
	}

	for (idx < lenS) && ('0' <= symb && symb <= '9') {
		symb = s[idx]
		res = res*10 + int(symb-'0')
		if sign*res > math.MaxInt32 || sign*res < math.MinInt32 {
			return math.MaxInt32
		}
		idx++
	}
	return res * sign
}*/
