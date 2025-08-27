package main

import "testing"

func Test_deleteNode(t *testing.T) {

	tests := []struct {
		name string
		node *ListNode
	}{
		{name: "test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteNode(tt.node)
		})
	}
}
