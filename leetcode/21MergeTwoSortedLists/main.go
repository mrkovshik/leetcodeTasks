package main

import "fmt"

//Definition for singly-linked list.
 type ListNode struct {
   Val int
   Next *ListNode
  }



func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	switch{
	case list1==nil && list2==nil: 
	return nil
	case list1==nil && list2!=nil:
		return list2
	case list1!=nil && list2==nil:
		return list1
	}
    if list1.Val<list2.Val {
		list1.Next=mergeTwoLists(list1.Next, list2)
		return list1
	} 
	list2.Next=mergeTwoLists(list2.Next, list1)
		return list2
	
}

 
func printList (n *ListNode, x *int){

	fmt.Printf("Node %v = %v    ", *x, n.Val)
	*x++
	if n.Next!=nil {
		printList(n.Next, x)
	}
}
func main(){
	var x *int
	x=new(int)

	L1_2:= ListNode {
		Val: 6,
		Next: nil,
	}
	L1_1:= ListNode {
		Val:5,
		Next: &L1_2,
	}
	Head1:= ListNode {
		Val: 1,
		Next: &L1_1,
	}

	L2_2:= ListNode {
		Val: 4,
		Next: nil,
	}
	L2_1:= ListNode {
		Val: 3,
		Next: &L2_2,
	}
	Head2:= ListNode {
		Val: 2,
		Next: &L2_1,
	}
	fmt.Printf("List 1:")
	fmt.Println()
	printList (&Head1,x)
	fmt.Println()
	fmt.Printf("List 2:")
	fmt.Println()
	*x=0
	printList (&Head2,x)

	newhead:=mergeTwoLists(&Head1,&Head2)
*x=0
fmt.Println()
fmt.Println("New List:")
	 printList(newhead,x)
	

	// *x=0
	// if Head1.Val>=Head2.Val{
	// 	printList(&Head2,x)
	// 	} else {
	// 		fmt.Println(&Head1,x)
	// 	}
}