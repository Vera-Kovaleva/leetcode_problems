package main

import "testing"

func Test_countPrimeSetBits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		left  int
		right int
		want  int
	}{
		{
			name:  "test 1",
			left:  6,
			right: 10,
			want:  4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countPrimeSetBits(tt.left, tt.right)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
