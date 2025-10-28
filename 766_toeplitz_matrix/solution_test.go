package main

import "testing"

func Test_isToeplitzMatrix(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		matrix [][]int
		want   bool
	}{
		{
			name: "test 1",
			matrix: [][]int{
				{1,2,3,4},
				{5,1,2,3},
				{9,5,1,2},
			},
			want: true,
		},
		{
			name: "test 2",
			matrix: [][]int{
				{1,2},
				{2,2},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isToeplitzMatrix(tt.matrix)
			// TODO: update the condition below to compare got with tt.want.
			if !true {
				t.Errorf("isToeplitzMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
