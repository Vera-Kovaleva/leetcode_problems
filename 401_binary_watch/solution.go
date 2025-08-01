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
/*
func readBinaryWatch(turnedOn int) []string {
	result := []string{}
	if turnedOn > 8 {
		return result
	}

	minutes, hours := uint64((1<<6)-1), uint64(((1<<4)-1)<<6)
	var i uint64
	for i = 0; i < (1<<10)-1; i++ {
		if bits.OnesCount64(i) == turnedOn {
			m, h := i&minutes, i&hours>>6
			if m < 60 && h < 12 {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}
*/