package main

import (
	"fmt"

)

func majorityElement(nums []int) int {
	halfLen := len(nums) / 2
	numsMap := make(map[int]int)
	for _, n := range nums {
		numsMap[n]++
		if numsMap[n] > halfLen {
			return n
		}
	}
	return 0
}

func main() {
	nums := []int{3}
	fmt.Print(majorityElement(nums))
}
