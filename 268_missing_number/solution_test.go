package main

import "testing"

func Test_missingNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{3,0,1},
			want: 2,
		},
		{
			name: "test 2",
			nums: []int{0,1},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := missingNumber(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if (got!=tt.want) {
				t.Errorf("missingNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
