package main

import "testing"

func Test_countTriples(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		n    int
		want int
	}{
		{
			name: "test 1",
			n: 5,
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countTriples(tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got!=tt.want {
				t.Errorf("countTriples() = %v, want %v", got, tt.want)
			}
		})
	}
}
