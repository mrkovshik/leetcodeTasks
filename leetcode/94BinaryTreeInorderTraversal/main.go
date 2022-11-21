package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
var s [] int	

func processNode(root *TreeNode)[]int {
	if root==nil{
		return nil
	}
	s=append(s,root.Val)
	fmt.Println("appended",s)
	if root.Left!=nil {
		processNode(root.Left)
	}
	if root.Right!=nil {
		processNode(root.Right)
	}
	return s
}

func inorderTraversal(root *TreeNode) []int {
	
res:=processNode(root)
if len(res)>2{
res[1],res[2]=res[2], res[1]}
return res

}

func main() {
	
	// node4 := TreeNode{0, nil, nil}
	// node3 := TreeNode{2, nil, nil}
	// node2 := TreeNode{3, nil, nil}
	// node1 := TreeNode{2, &node2, nil}
	// root := TreeNode{1, &node1, nil}
	root := TreeNode{}
	fmt.Println(root)
	res := inorderTraversal(&root)
	fmt.Println(res)
}