package main

func romanToInt(s string) int {
	match := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	res := 0
	last := 0

	for i := len(s) - 1; i >= 0; i-- {
		curr := match[s[i]]
		if curr >= last {
			res += curr
			last = curr
		} else {
			res -= curr
		}
	}

	return res
}
