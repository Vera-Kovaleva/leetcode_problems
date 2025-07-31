package main

import (
	"fmt"
	"math/bits"
)

//https://leetcode.com/problems/binary-watch/description/

func readBinaryWatch(turnedOn int) []string {
	var result []string

	for hour := 0; hour < 12; hour++ {
		for min := 0; min < 60; min++ {
			if bits.OnesCount(uint(hour))+bits.OnesCount(uint(min)) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", hour, min))
			}
		}
	}

	return result
}
