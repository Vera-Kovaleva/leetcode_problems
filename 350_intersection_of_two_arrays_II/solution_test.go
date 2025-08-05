package main

import (
	"reflect"
	"testing"
)

func Test_intersect(t *testing.T) {

	tests := []struct {
		name string
		nums1 []int
		nums2 []int
		want []int
	}{
		{
			name: "test 1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			want: []int{2, 2},
		},
		{
			name: "test 2",
			nums1: []int{4, 9, 5},
			nums2: []int{9,4,9,8,4},
			want: []int{4, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersect(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
