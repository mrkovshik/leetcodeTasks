package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
func isSameTree(p *TreeNode, q *TreeNode) bool {
if p!=nil && q!=nil {
	if p.Val!=q.Val{
		return false
	}
	if isSameTree(p.Right,q.Right)&& isSameTree(p.Left,q.Left){
		return true
	} 
	}
	if p==nil&& q==nil {
		return true
	}

	return false
}
func main() {
	node13 := TreeNode{3, nil, nil}
	node11 := TreeNode{1, &node13, nil}
	node12 := TreeNode{2, &node13, nil}
	node23 := TreeNode{3, nil, nil}
	node21 := TreeNode{1, &node23, nil}
	node22 := TreeNode{2, nil, nil}
	p := TreeNode{0, &node11, &node12}
	q := TreeNode{0, &node21, &node22}
if isSameTree(&p,&q){
	fmt.Printf("Trees are same")
} else {
	fmt.Printf("Trees are not same")
}
}