package main

func isPalindrome(x int) bool {
	frontWards := x
	backWards := 0
	for x > 0 {
		backWards = backWards*10 + x%10
		x /= 10
	}

	return backWards==frontWards
}
