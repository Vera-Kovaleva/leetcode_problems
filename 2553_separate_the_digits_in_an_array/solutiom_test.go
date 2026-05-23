package main

import "testing"

func slice_the_same(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func Test_separateDigits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want []int
	}{
		{
			name: "test1",
			nums: []int{13, 25, 83, 77},
			want: []int{1, 3, 2, 5, 8, 3, 7, 7},
		},
		{
			name: "test2",
			nums: []int{7, 1, 3, 9},
			want: []int{7, 1, 3, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := separateDigits(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if !slice_the_same(tt.want, got) {
				t.Errorf("separateDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
