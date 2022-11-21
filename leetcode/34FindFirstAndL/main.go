package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	min := -1
	max := -1
	found := false
	res := make ([]int,2)
	for i, j := range nums {
		fmt.Println("nums[",i,"]=",j)
		if j == target {
			if found == false {
				min = i
				max=i
				found = true
			} else {
				max = i
			}

		}
		fmt.Println("min=", min, "max=", max)
	}
	if min != -1 && max == -1 {
		max = len(nums) - 1
	}
	res[0]=min
	res[1]=max
	return res
}

func main() {
	target := 3
	nums := []int{1, 2, 3, 3, 3, 2}
	fmt.Println("nums=", nums, "target=", target)
	fmt.Println(searchRange(nums, target))
}
