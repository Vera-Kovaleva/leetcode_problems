package main

import "testing"

func Test_modifiedList(t *testing.T) {
	createList := func(vals []int) *ListNode {
		if len(vals) == 0 {
			return nil
		}
		head := &ListNode{Val: vals[0]}
		current := head
		for i := 1; i < len(vals); i++ {
			current.Next = &ListNode{Val: vals[i]}
			current = current.Next
		}
		return head
	}
	listToSlice := func(head *ListNode) []int {
		var result []int
		curr := head
		for curr != nil {
			result = append(result, curr.Val)
			curr = curr.Next
		}
		return result
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		head *ListNode
		want *ListNode
	}{
		{
			name: "test 1",
			nums: []int{1, 2, 3},
			head: createList([]int{1, 2, 3, 4, 5}),
			want: createList([]int{4, 5}),
		},
		{
			name: "test 2",
			nums: []int{1},
			head: createList([]int{1,2,1,2,1,2}),
			want: createList([]int{2, 2, 2}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := modifiedList(tt.nums, tt.head)

			gotSlice := listToSlice(got)
			wantSlice := listToSlice(tt.want)

			if len(gotSlice) != len(wantSlice) {
				t.Errorf("modifiedList() = %v, want %v", gotSlice, wantSlice)
				return
			}

			for i := range gotSlice {
				if gotSlice[i] != wantSlice[i] {
					t.Errorf("modifiedList() = %v, want %v", gotSlice, wantSlice)
					return
				}
			}
		})
	}
}
