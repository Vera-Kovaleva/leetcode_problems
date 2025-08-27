package main

func reverse(x int) int {
	res, syng := 0, 1
	if x < 0 {
		syng = -1
		x = x * -1
	}

	for x > 0 {
		res = res*10 + x%10
		x = x / 10
	}

	if res > 2147483647 {
		return 0
	}

	return res * syng
}
