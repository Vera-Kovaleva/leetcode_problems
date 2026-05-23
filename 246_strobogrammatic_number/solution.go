package main

func isStrobogrammatic(num string) bool {
	rotatedDigits := map[byte]byte{'0': '0', '1': '1', '8': '8', '6': '9', '9': '6'}
	for i, j := 0, len(num)-1; i <= j; i, j = i+1, j-1 {
		rotated, ok := rotatedDigits[num[i]]
		if ok != true || rotated != num[j] {
			return false
		}
	}
	return true
}
