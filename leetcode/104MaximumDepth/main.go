package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {

}
func main() {
	node6 := TreeNode{2, nil, nil}
	node5 := TreeNode{5, nil, &node6}
	node4 := TreeNode{5, nil, nil}
	node3 := TreeNode{2, nil, nil}
	node2 := TreeNode{1, &node3, &node5}
	node1 := TreeNode{1, nil, &node4}
	head := TreeNode{0, &node1, &node2}
	fmt.Printf("the tree is %v", maxDepth(&head))
}