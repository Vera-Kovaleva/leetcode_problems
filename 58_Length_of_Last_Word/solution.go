package main

/*
func lengthOfLastWord(s string) int {
	var result = 0
	var flag = 0
	for _, symb := range s {
		if symb != ' ' {
			if flag == 1 {
				flag = 0
				result = 0
			}
			result += 1
		} else {
			flag = 1
		}
	}
	return result
}
*/

func lengthOfLastWord(s string) int {
	var result = 0
	var n = len(s)
	for i := n - 1; i >= 0; i-- {
		if s[i] != ' ' {
			result += 1
		} else if result > 0 {
			break
		}
	}
	return result
}
