package main

import "fmt"

//Definition for singly-linked list.
 type ListNode struct {
   Val int
    Next *ListNode
}

func swap (node *ListNode){
	if node.Next !=nil {
		node.Val, node.Next.Val= node.Next.Val, node.Val
		if node.Next.Next != nil {
			swap(node.Next.Next)
		}
	}
	return
}
 func swapPairs(head *ListNode) *ListNode {
	if head!=nil{
    swap(head)
	}
return head
 }


func main (){
fifNode:=	ListNode {5, nil}
forNode:= ListNode {4, &fifNode}
thrdNode:= ListNode {3, &forNode}
secNode:= ListNode {2, &thrdNode}
head:= ListNode {1, &secNode}

fmt.Printf("head = %v\n secNode = %v \n thrdNode = %v \n forNode = %v \n fifNode = %v \n ",head, secNode,thrdNode,forNode, fifNode )
swapPairs(&head)
fmt.Printf("head = %v\n secNode = %v \n thrdNode = %v \n forNode = %v \n fifNode = %v \n ",head, secNode,thrdNode,forNode, fifNode )

}



