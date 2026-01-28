package main

import "testing"

func Test_kLengthApart(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want bool
	}{
		{
			name: "test 2",
			nums: []int{1,0,0,1,0,1},
			k: 2,
			want: false,
		},
		{
			name: "test 1",
			nums: []int{1,0,0,0,1,0,0,1},
			k: 2,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kLengthApart(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("kLengthApart() = %v, want %v", got, tt.want)
			}
		})
	}
}
