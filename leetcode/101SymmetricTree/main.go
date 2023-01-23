package main

import (
	"fmt"

	)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mirr (nodeL *TreeNode, nodeR *TreeNode)bool {
	var outer, inner bool
if nodeL==nil && nodeR == nil {
	return true
}
if nodeL==nil || nodeR == nil {
	return false
}

switch{
case nodeL.Left==nil&& nodeR.Right==nil:
outer=true
case nodeL.Left!=nil&& nodeR.Right!=nil:
	if nodeL.Left.Val==nodeR.Right.Val{
outer=mirr(nodeL.Left, nodeR.Right)
	} else {
		return false
	}
default: return false
}
switch{
case nodeR.Left==nil&& nodeL.Right==nil:
inner=true
case nodeL.Right!=nil&& nodeR.Left!=nil:
	if nodeL.Right.Val==nodeR.Left.Val{
inner=mirr(nodeL.Right, nodeR.Left)
	} else {
		return false
	}
default: return false
}
if inner&&outer{
	return true
} else {
	return false
}
}

func isSymmetric(root *TreeNode) bool {
    if root == nil {
		return true
	}
	if root.Left==nil&& root.Right==nil{
		return true
	}
	if root.Left==nil|| root.Right==nil{
		return false
	}
	if root.Left.Val== root.Right.Val{
		return mirr(root.Left, root.Right)
	}
	return false
}

func main() {
	// node6 := TreeNode{2, nil, nil}
	node5 := TreeNode{5, nil, nil}
	node4 := TreeNode{5, nil, nil}
	// node3 := TreeNode{2, nil, nil}
	node2 := TreeNode{1, &node5, nil}
	node1 := TreeNode{1,nil, &node4}
	head := TreeNode{0, &node1, &node2}
	fmt.Printf("Answer is %v", isSymmetric(&head))
}