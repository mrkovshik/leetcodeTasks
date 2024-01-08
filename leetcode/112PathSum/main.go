package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	return calculateNode(0, targetSum, *root)
}

func calculateNode(current, targetSum int, node TreeNode) bool {
	var left, right = false, false
	current += node.Val
	if current > targetSum {
		return false
	}
	if node.Left == nil && node.Right == nil {
		if current == targetSum {
			return true
		}
		return false
	}

	if node.Left != nil {
		left = calculateNode(current, targetSum, *node.Left)
	}
	if node.Right != nil {
		right = calculateNode(current, targetSum, *node.Right)
	}
	return left || right
}

func main() {
	targetSum := 5
	// node8 := TreeNode{1, nil, nil}
	// node7 := TreeNode{3, nil, nil}
	// node6 := TreeNode{7, nil, nil}
	// node5 := TreeNode{5, nil, &node8}
	// node4 := TreeNode{5, nil, nil}
	// node3 := TreeNode{11, &node6, &node7}
	node2 := TreeNode{8, nil, nil}
	node1 := TreeNode{4, nil, nil}
	head := TreeNode{5, &node1, &node2}
	fmt.Printf("the tree is %v", hasPathSum(&head, targetSum))
}
