package main

import (
	"container/heap"
	"fmt"
	
)
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	x := old[len(old) - 1]
	*h = old[0 : len(old) - 1]
	return x
}

func minStoneSum(piles []int, k int) int {
    h := &IntHeap{}
    heap.Init(h)
    sum := 0
    
    for _, pile := range piles {
        heap.Push(h, pile)
        sum += pile
    }
    
    for i := 0; i < k; i++ {
        temp := heap.Pop(h).(int)
        sum -= (temp / 2)
        heap.Push(h, temp - (temp / 2))
    }
    
    return sum
}

func main() {
	piles := []int{7481,7973,3635,5320,2721,4394,7861}
	k := 10
	fmt.Printf("there is %v stones minimum remain", minStoneSum(piles, k))
}