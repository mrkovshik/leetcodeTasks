package main

import (
	"fmt"
	"strconv"
)

var mltplr *int
var res *int

type ListNode struct {
	Val  int
	Next *ListNode
}

func convertListToInt(n *ListNode) int {
	if n != nil {
		*res = *res + n.Val**mltplr
		*mltplr *= 10
		*res = convertListToInt(n.Next)
	}
	return *res
}
func convertIntToList(n int) *ListNode {
	dcm := 10
	listLen := len(strconv.Itoa(n))
	nodes := make([]ListNode, listLen)
	if listLen == 1 {
		nodes[0].Val = n
		nodes[0].Next = nil
	}
	for i := 0; i < listLen; i++ {
		if i == listLen-1 {
			// fmt.Println("DCM= ", dcm)
			// fmt.Println("n= ", n)
			// fmt.Println("val= ", n%dcm)
			nodes[i].Val = n % dcm / (dcm / 10)
			nodes[i].Next = nil
		} else {
			// fmt.Println("DCM= ", dcm)
			// fmt.Println("n= ", n)
			// fmt.Println("val= ", n%dcm)
			nodes[i].Val = n % dcm / (dcm / 10)
			nodes[i].Next = &nodes[i+1]
			n = n - n%dcm
			dcm *= 10
		}
	}
	return &nodes[0]
}

func printList(n *ListNode, x *int) {
	fmt.Printf("Node %v = %v    ", *x, n.Val)
	*x++
	if n.Next != nil {
		printList(n.Next, x)
	}
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode
	newList = new(ListNode)
	mltplr = new(int)
	*mltplr = 1
	res = new(int)
	Num1 := convertListToInt(l1)
	fmt.Printf("Num1=%v", Num1)
	*res = 0
	*mltplr = 1
	Num2 := convertListToInt(l2)
	fmt.Printf("Num2=%v", Num2)
	Num3 := Num1 + Num2
	fmt.Printf("Num3=%v", Num3)
	newList = convertIntToList(Num3)
	return newList
}

func main() {

	var x *int
	x = new(int)
	L1_3 := ListNode{
		Val:  3,
		Next: nil,
	}

	L1_2 := ListNode{
		Val:  3,
		Next: &L1_3,
	}
	L1_1 := ListNode{
		Val:  4,
		Next: &L1_2,
	}
	Head1 := ListNode{
		Val:  2,
		Next: &L1_1,
	}

	L2_2:= ListNode {
		Val: 4,
		Next: nil,
	}
	L2_1:= ListNode {
		Val: 5,
		Next: &L2_2,
	}
	Head2 := ListNode{
		Val:  5,
		Next: &L2_1,
	}
	Head3 := addTwoNumbers(&Head1, &Head2)

	fmt.Printf("List 1:")
	fmt.Println()
	printList(&Head1, x)
	fmt.Println()
	fmt.Printf("List 2:")
	fmt.Println()
	*x = 0
	printList(&Head2, x)
	fmt.Println()
	fmt.Printf("List 3:")
	fmt.Println()
	*x = 0
	printList(Head3, x)
}
