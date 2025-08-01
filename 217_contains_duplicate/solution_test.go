package main

import "testing"

func Test_containsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "test 1",
			nums: []int{1,2,3,1},
			want: true,
		},
		{
			name: "test 1",
			nums: []int{1,1},
			want: true,
		},
		{
			name: "test 1",
			nums: []int{1,2,3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsDuplicate(tt.nums); got != tt.want {
				t.Errorf("containsDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
