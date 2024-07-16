package main

import (
	"fmt"
)

// this is a comment

//Definition for a binary tree node.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getMinimumDifference(root *TreeNode) int {
	const INF = 999999999
	res, last := INF, INF

	var minDiff func(*TreeNode)
	minDiff = func(root *TreeNode) {
		if root == nil {
			return
		}
		minDiff(root.Left)
		res = min(res, abs(last-root.Val))
		last = root.Val
		minDiff(root.Right)
	}

	minDiff(root)
	return res
}

func main() {
	fmt.Print("yes")
}
