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

func lengthOfLastWord(s string) (result int) {
	for i := len(s) - 1; i >= 0 && (s[i] != ' ' || result == 0); i-- {
		if s[i] != ' ' {
			result++
		}
	}

	return
}
