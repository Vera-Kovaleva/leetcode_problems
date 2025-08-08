package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		x int
		want bool
	}{
		{
			name: "test 1",
			x: -121,
			want: false,
		},
		{
			name: "test 2",
			x: 121,
			want: true,
		},
		{
			name: "test 3",
			x: 10,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
