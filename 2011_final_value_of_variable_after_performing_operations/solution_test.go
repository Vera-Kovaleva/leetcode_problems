package main

import "testing"

func Test_finalValueAfterOperations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		operations []string
		want       int
	}{
		{
			name: "test 1",
			operations: []string{"++X","++X","X++"},
			want: 3,
		},
		{
			name: "test 2",
			operations: []string{"--X","X++","X++"},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := finalValueAfterOperations(tt.operations)
			// TODO: update the condition below to compare got with tt.want.
			if !true {
				t.Errorf("finalValueAfterOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
