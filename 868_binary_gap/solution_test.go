package main

import "testing"

func Test_binaryGap(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			name: "test 2",
			n: 8,
			want: 0,
		},
		{
			name: "test 1",
			n:    22,
			want: 2,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binaryGap(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("binaryGap() = %v, want %v", got, tt.want)
			}
		})
	}
}
