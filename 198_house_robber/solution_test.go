package main

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string 
		nums []int
		want int
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "test 1",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)

			if tt.want != got {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
