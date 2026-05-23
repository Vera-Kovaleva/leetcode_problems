package main

func checkOnesSegment(s string) bool {
	n := len(s)
	counter := 0

	for i := 1; i < n; i++ {
		if s[i] == '0' {
			counter++
		}
		if counter > 0 && s[i] == '1' {
			return false
		}
	}
	return true
}

/*
func checkOnesSegment(s string) bool {
	return !strings.Contains(s, "01")
}
	*/