package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	level := 0
	nodeQueue := []*TreeNode{root}

	for len(nodeQueue) > 0 {
		for _, node := range nodeQueue {
			nodeQueue = nodeQueue[1:]
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		level++
	}
	return level
}
