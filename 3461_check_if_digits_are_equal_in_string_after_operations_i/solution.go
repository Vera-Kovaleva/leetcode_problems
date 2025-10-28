package main

func hasSameDigits(s string) bool {
	if len(s) == 2 {
		if s[0] == s[1] {
			return true
		}
		return false
	}

	curS := ""
	for len(s) > 2 {
		for i := 0; i < len(s); i++ {
			curS += string('a')
		}
	}

	return false
}
