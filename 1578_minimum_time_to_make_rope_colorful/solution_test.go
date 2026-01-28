package main

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		colors     string
		neededTime []int
		want       int
	}{
		{
			name:       "test 1",
			colors:     "aaabbbabbbb",
			neededTime: []int{3,5,10,7,5,3,5,5,4,8,1},
			want:       26,
		},
		{
			name:       "test 2",
			colors:     "abc",
			neededTime: []int{1, 2, 3},
			want:       0,
		},
		{
			name:       "test 3",
			colors:     "aabaa",
			neededTime: []int{1,2,3,4,1},
			want:       2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minCost(tt.colors, tt.neededTime)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
