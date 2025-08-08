package main

import "testing"

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "test 1",
			s:    "aabbcc",
			want: "abccba",
		},
		{
			name: "test 2",
			s:    "rat",
			want: "art",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortString(tt.s); got != tt.want {
				t.Errorf("sortString() = %v, want %v", got, tt.want)
			}
		})
	}
}
