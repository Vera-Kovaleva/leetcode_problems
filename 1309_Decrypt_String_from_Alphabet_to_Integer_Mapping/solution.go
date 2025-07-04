package main

/*
func freqAlphabets(s string) string {
	dictionary := make(map[string]string)

	dictionary["1"] = "a"
	dictionary["2"] = "b"
	dictionary["3"] = "c"
	dictionary["4"] = "d"
	dictionary["5"] = "e"
	dictionary["6"] = "f"
	dictionary["7"] = "g"
	dictionary["8"] = "h"
	dictionary["9"] = "i"
	dictionary["10"] = "j"
	dictionary["11"] = "k"
	dictionary["12"] = "l"
	dictionary["13"] = "m"
	dictionary["14"] = "n"
	dictionary["15"] = "o"
	dictionary["16"] = "p"
	dictionary["17"] = "q"
	dictionary["18"] = "r"
	dictionary["19"] = "s"
	dictionary["20"] = "t"
	dictionary["21"] = "u"
	dictionary["22"] = "v"
	dictionary["23"] = "w"
	dictionary["24"] = "x"
	dictionary["25"] = "y"
	dictionary["26"] = "z"

	result := ""

	for i := len(s) - 1; i >= 0; i-- {
		number := ""
		if s[i] == '#' {
			i -= 2
			number = string(s[i]) + string(s[i+1])
		} else {
			number = string(s[i])
		}
		fmt.Println(number)
		result = dictionary[number] + result
	}

	return result
}
*/

import "strings"

//https://leetcode.com/problems/decrypt-string-from-alphabet-to-integer-mapping/

func freqAlphabets(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		n := s[i] - '0'
		if i+2 < len(s) && s[i+2] == '#' {
			n = 10*n + s[i+1] - '0'
			i += 2
		}
		result.WriteByte('a' + n - 1)
	}
	return result.String()
}