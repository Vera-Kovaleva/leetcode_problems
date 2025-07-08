package main

//https://leetcode.com/problems/crawler-log-folder/

func minOperations(logs []string) int {
	ans := 0
	for _, log := range logs {
		if log == "../" {
			if ans > 0 {
				ans--
			}
		} else if log != "./" {
			ans++
		}
	}
	return ans
}
