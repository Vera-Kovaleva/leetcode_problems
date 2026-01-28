package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}
*/

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	sumStack := []int{}
	nodeStack := []*TreeNode{}
	sumStack = append(sumStack, sum-root.Val)
	nodeStack = append(nodeStack, root)
	for len(nodeStack) > 0 {
		lastIndex := len(nodeStack) - 1
		currNode := nodeStack[lastIndex]
		nodeStack = nodeStack[:lastIndex]
		currSum := sumStack[lastIndex]
		sumStack = sumStack[:lastIndex]
		if currNode.Left == nil && currNode.Right == nil && currSum == 0 {
			return true
		}
		if currNode.Left != nil {
			nodeStack = append(nodeStack, currNode.Left)
			sumStack = append(sumStack, currSum-currNode.Left.Val)
		}
		if currNode.Right != nil {
			nodeStack = append(nodeStack, currNode.Right)
			sumStack = append(sumStack, currSum-currNode.Right.Val)
		}
	}
	return false
}
