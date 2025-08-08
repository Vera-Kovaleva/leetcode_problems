package main

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{
			name: "test 1",
			s:    " ",
			want: true,
		},
		{
			name: "test 2",
			s:    "A man, a plan, a canal: Panama",
			want: true,
		},
		{
			name: "test 3",
			s:    "race a car",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.s); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
