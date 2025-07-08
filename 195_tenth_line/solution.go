package main

//https://leetcode.com/problems/tenth-line/description/

import (
	"bufio"
	"os"
)

func temLines(fileName string) string {
	result := ""

	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}

	scanner := bufio.NewScanner(file)

	counter := 0
	for scanner.Scan() {
		if counter == 10 {
			return result
		}
		if counter != 0 {
			result += "\n"
		}
		line := scanner.Text()
		result += line
		counter++
	}

	return result
}
