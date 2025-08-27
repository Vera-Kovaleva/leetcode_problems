package main

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "test 1",
			x:    321,
			want: 123,
		},
		{
			name: "test 1",
			x:    -123,
			want: -321,
		},
		{
			name: "test 1",
			x:    120,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
