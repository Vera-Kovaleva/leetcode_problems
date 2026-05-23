package main

import "testing"

func Test_isStrobogrammatic(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		num  string
		want bool
	}{
		{
			name: "test 4",
			num: "818",
			want: true,
		},
		{
			name: "test 1",
			num: "88",
			want: true,
		},
		{
			name: "test 2",``
			num: "69",
			want: true,
		},
		{
			name: "test 3",
			num: "8894",
			want: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isStrobogrammatic(tt.num)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want!=got {
				t.Errorf("isStrobogrammatic() = %v, want %v", got, tt.want)
			}
		})
	}
}
