package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {

	counter := make(map[int]bool, len(nums))
	for _, num := range nums {
		counter[num] = true
	}

	start := &ListNode{Next: head}
	cur := start

	for cur.Next != nil {
		if counter[cur.Next.Val] {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}

	}
	return start.Next
}
