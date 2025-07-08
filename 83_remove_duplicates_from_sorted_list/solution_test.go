package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{{
		name: "empty list",
		args: args{head: nil},
		want: nil,
	},
		{
			name: "single node",
			args: args{head: &ListNode{Val: 1}},
			want: &ListNode{Val: 1},
		},
		{
			name: "one duplicate",
			args: args{head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2}}}},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
