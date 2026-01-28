package main

import "testing"

func Test_getSneakyNumbers(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "test 1",
			nums: []int{0, 1, 1, 0},
			want: []int{0, 1},
		},
		{
			name: "test 2",
			nums: []int{0,3,2,1,3,2},
			want: []int{2, 3},
		},
		{
			name: "test 3",
			nums: []int{7,1,5,4,3,4,6,0,9,5,8,2},
			want: []int{4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSneakyNumbers(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("getSneakyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
