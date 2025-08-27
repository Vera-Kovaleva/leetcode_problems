package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	res := &ListNode{0, head}
	first, second := res, res

	for i := 0; i <= n; i++ {
		first = first.Next
	}
	if first == nil {
		return second
	}
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next

	return res.Next
}
