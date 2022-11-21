package main

import(
	"fmt"
)
  
  type ListNode struct {
  Val int
  Next *ListNode
}

 func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
		} else
		if head.Next != nil {
	if head.Val==head.Next.Val{
			head=deleteDuplicates(head.Next)	
				} else {
	head.Next=deleteDuplicates( head.Next)	
			}
		}
		return head
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

	L4:= ListNode {
		Val: 3,
		Next: nil,
	}
	L3:= ListNode {
		Val: 2,
		Next: &L4,
	}
	L2:= ListNode {
		Val: 2,
		Next: &L3,
	}
	L1:= ListNode {
		Val:2,
		Next: &L2,
	}
	Head:= ListNode {
		Val: 1,
		Next: &L1,
	}

	fmt.Printf("List:")
	fmt.Println()
	printList (&Head,x)
	fmt.Println()

	
newhead:=deleteDuplicates(&Head)
*x=0
fmt.Println()
fmt.Println("New List:")
	 printList(newhead,x)
	}