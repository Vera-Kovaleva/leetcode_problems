package main

import "testing"

func Test_myAtoi(t *testing.T) {
	tests := []struct {
		name string
		s string
		want int
	}{
		{
			name: "tets 1",
			s: "42",
			want: 42,
		},
		{
			name: "tets 1",
			s: "  -42",
			want: -42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
