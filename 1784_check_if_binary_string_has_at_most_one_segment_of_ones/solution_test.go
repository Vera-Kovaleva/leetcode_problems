package main

import "testing"

func Test_checkOnesSegment(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want bool
	}{

		{
			name: "test 1",
			s:    "1001",
			want: false,
		},
		{
			name: "test 2",
			s:    "1100",
			want: true,
		},
		{
			name: "test 3",
			s:    "10011",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := checkOnesSegment(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("checkOnesSegment() = %v, want %v", got, tt.want)
			}
		})
	}
}
