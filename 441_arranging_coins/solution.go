package main

func arrangeCoins(n int) int {
	i := 0
	for n>=i+1 {
		i++
		n-=i
	}
	return i
}
