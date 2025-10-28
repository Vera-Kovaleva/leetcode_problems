package main

import "testing"

func Test_findMissingRanges(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums  []int
		lower int
		upper int
		want  [][]int
	}{
		{
			name: "test 1",
			nums: []int{0,1,3,50,75},
			lower: 0,
			upper: 99,
			want: [][]int{{2,2},{4,49},{51,74},{76,99}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMissingRanges(tt.nums, tt.lower, tt.upper)
			// TODO: update the condition below to compare got with tt.want.
			if !true {
				t.Errorf("findMissingRanges() = %v, want %v", got, tt.want)
			}
		})
	}
}
