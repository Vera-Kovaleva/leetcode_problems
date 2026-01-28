package main

import "sort"

func isValidCode(code string) bool {
	if len(code) == 0 {
		return false
	}
	for _, r := range code {
		if !((r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_') {
			return false
		}
	}
	return true
}

func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	result := make([]string, 0, len(code))
	buisnesses := []string{"electronics", "grocery", "pharmacy", "restaurant"}
	l, r := 0, 0
	for _, buisness := range buisnesses {
		for i := 0; i < len(code); i++ {
			if isActive[i] && businessLine[i] == buisness && isValidCode(code[i]) {
				result = append(result, code[i])
				r++
			}
		}
		sort.Strings(result[l:r])
		l = r
	}
	return result
}
