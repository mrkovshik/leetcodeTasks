package main
import(
	"fmt"
)
  
  type ListNode struct {
      Val int
     Next *ListNode
  }
func removal (dub int, node *ListNode) *ListNode {
		if node.Val==dub {
		if node.Next==nil{
node=nil
} else {
	node=node.Next
	node=removal(dub,node)
		} 
	}
	 return node
	}
	 
 func deleteDuplicates(head *ListNode) *ListNode {
    if head==nil {
		return nil
	}
	if head.Next!=nil{
	if head.Val==head.Next.Val{
		head=removal(head.Val,head)
		head=deleteDuplicates(head)
	}else {
			head.Next=deleteDuplicates(head.Next)
		}
		} 
		return head
	} 	
	
func printList (n *ListNode, x *int){
if n==nil {
	fmt.Println("nothing left")
} else {
	fmt.Printf("Node %v = %v    ", *x, n.Val)
	*x++
	if n.Next!=nil {
		printList(n.Next, x)
	}
}
}
func main (){
	
	var x *int
	x=new(int)
	
	L6:= ListNode {
		Val: 2,
		Next: nil,
	}
	L5:= ListNode {
		Val: 2,
		Next: &L6,
	}
	L4:= ListNode {
		Val: 2,
		Next: &L5,
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
		Val:1,
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
