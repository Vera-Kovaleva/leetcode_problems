package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const null = -1

func buildTree(rootIndex int, values []int) *TreeNode {
	var node, leftNode, rightNode *TreeNode

	leftIndex, rightIndex := 2*rootIndex+1, 2*rootIndex+2
	if leftIndex < len(values) && values[leftIndex] != null {
		leftNode = buildTree(leftIndex, values)
	}
	if rightIndex < len(values) && values[rightIndex] != null {
		rightNode = buildTree(rightIndex, values)
	}
	if rootIndex < len(values) {
		node = &TreeNode{
			Val:   values[rootIndex],
			Left:  leftNode,
			Right: rightNode,
		}
	}

	return node
}
