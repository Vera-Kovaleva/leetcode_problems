package main

import "testing"

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		columnNumber int
		want         string
	}{
		{
			name:         "test 3",
			columnNumber: 701,
			want:         "ZY",
		},
		{
			name:         "test 1",
			columnNumber: 1,
			want:         "A",
		},
		{
			name:         "test 2",
			columnNumber: 28,
			want:         "AB",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := convertToTitle(tt.columnNumber)
			// TODO: update the condition below to compare got with tt.want.
			if tt.want != got {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
