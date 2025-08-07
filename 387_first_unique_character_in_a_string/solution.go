package main

func firstUniqChar(s string) int {
	counter := make([]int, 26)
	for _, symb := range s {
		counter[symb-'a']++
	}
	for i, symb := range s {
		if counter[symb-'a'] == 1 {
			return i
		}
	}
	return -1
}

/*
func firstUniqChar(s string) int {
	counter := make([]int, 29)
	stopNumber := 100001
	for i, symb := range s {
		curSymbNumber := int(symb - 'a')
		if counter[curSymbNumber] != 0 {
			counter[curSymbNumber] = stopNumber
		} else {
			counter[curSymbNumber] = i + 1
		}
	}
	ans := stopNumber
	for _, index := range counter {
		if index != 0 && index != stopNumber {
			ans = min(index-1, ans)
		}
	}
	if ans == stopNumber {
		ans = -1
	}
	return ans
}
*/
