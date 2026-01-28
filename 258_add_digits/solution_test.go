package main

import "testing"

func Test_addDigits(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  int
		want int
	}{
		{
			name: "test 1",
			num: 38,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addDigits(tt.num)
			// TODO: update the condition below to compare got with tt.want.
			if got!=tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
