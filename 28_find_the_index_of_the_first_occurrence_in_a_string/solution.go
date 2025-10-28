package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	hLen, nLen := len(haystack), len(needle)
	for i := 0; i < hLen-nLen+1; i++ {
		if haystack[i:i+nLen] == needle {
			return i
		}
	}
	return -1
}
