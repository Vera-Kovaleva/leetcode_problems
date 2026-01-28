package main

import (
	"testing"
)

func theSame(slice1 []string, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i]!=slice2[i] {
			return false
		}
	}
	return true
}

func Test_generatePossibleNextMoves(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		currentState string
		want         []string
	}{
		{
			name: "test 1",
			currentState: "++++",
			want: []string{"--++","+--+","++--"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generatePossibleNextMoves(tt.currentState)
			// TODO: update the condition below to compare got with tt.want.
			if !theSame(tt.want, got) {
				t.Errorf("generatePossibleNextMoves() = %v, want %v", got, tt.want)
			}
		})
	}
}
