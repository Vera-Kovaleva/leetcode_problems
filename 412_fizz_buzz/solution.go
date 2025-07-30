package main

import "strconv"

func fizzBuzzHelp(n int) string {
	res := ""
	if n%3 == 0 {
		res += "Fizz"
	}
	if n%5 == 0 {
		res += "Buzz"
	}
	if len(res) == 0 {
		return strconv.Itoa(n)
	}
	return res
}

func fizzBuzz(n int) []string {
	var res []string
	for i := 1; i <= n; i++ {
		//res = append(res, fizzBuzzHelp(i))
		curRes := ""
		if i%3 == 0 {
			curRes += "Fizz"
		}
		if i%5 == 0 {
			curRes += "Buzz"
		}
		if len(curRes) == 0 {
			curRes = strconv.Itoa(i)
		}
		res = append(res, curRes)
	}
	return res
}
