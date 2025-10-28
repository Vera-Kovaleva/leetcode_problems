package main

import "testing"

func Test_countValidSelections(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{1,0,2,0,3},
			want: 2,
		},
		{
			name: "test 2",
			nums: []int{2,3,4,0,4,1,0},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countValidSelections(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("countValidSelections() = %v, want %v", got, tt.want)
			}
		})
	}
}
