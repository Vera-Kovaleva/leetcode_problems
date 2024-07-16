package main

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{name: "Example 1", root: buildTree(0, []int{4, 2, 6, 1, 3}), want: 1},
		{name: "Example 2", root: buildTree(0, []int{1, 0, 48, null, null, 12, 49}), want: 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := getMinimumDifference(test.root); got != test.want {
				t.Errorf("getMinimumDifference() = %v, want %v", got, test.want)
			}
		})
	}
}
