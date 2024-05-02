package main

import "fmt"

func main() {
	nums := []int{-1, 10, 6, 7, -7, 1}
	fmt.Println(findMaxK(nums))
}

func findMaxK(nums []int) int {
	numsMapPos := make(map[int]struct{})
	numsMapNeg := make(map[int]struct{})
	var maxVal int
	for _, val := range nums {
		if val > 0 {
			if _, ok := numsMapPos[val]; ok {
				continue
			}
			numsMapPos[val] = struct{}{}
			if _, ok := numsMapNeg[-val]; ok {
				if val > maxVal {
					maxVal = val
				}
			}
			continue
		}
		if _, ok := numsMapNeg[val]; ok {
			continue
		}
		numsMapNeg[val] = struct{}{}
		if _, ok := numsMapPos[-val]; ok {
			if -val > maxVal {
				maxVal = -val
			}
		}

	}
	if maxVal > 0 {
		return maxVal
	}
	return -1
}
