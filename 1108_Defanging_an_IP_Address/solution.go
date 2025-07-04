package main

//https://leetcode.com/problems/defanging-an-ip-address/

import "strings"

func defangIPaddr(address string) string {
	var builder strings.Builder

	for i := 0; i < len(address); i++ {
		if address[i] == '.' {
			builder.WriteString("[.]")
		} else {
			builder.WriteByte(address[i])
		}
	}
	return builder.String()
}

/* 
конечно на литкоде есть такое решение, но как-то не спортивно
func defangIPaddr(address string) string {
    return strings.ReplaceAll(address,".","[.]")
}
*/