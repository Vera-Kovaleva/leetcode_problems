package main

import "strings"

func sortString(s string) string {
	var res strings.Builder

	counter := make([]int, 26)
	for _, symb := range s {
		counter[symb-'a']++
	}
	for res.Len() < len(s) {
		for i := 0; i <= 25; i++ {
			if counter[i] > 0 {
				res.WriteByte(byte(i + 'a'))
				counter[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if counter[i] > 0 {
				res.WriteByte(byte(i + 'a'))
				counter[i]--
			}
		}
	}
	return res.String()
}
