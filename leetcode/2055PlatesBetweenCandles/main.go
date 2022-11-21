package main

import (
	"fmt"
)

func platesBetweenCandles(s string, queries [][]int) []int {
var cnt int
	res := make([]int, len(queries))
	for i := range queries {
			cnt=0
			startFound:=false
		for j := queries[i][0]; j <=queries[i][1]; j++ {
			if s[j] == 124&& startFound==false{
				startFound=true
				}
				if startFound==true{
					if s[j]==42 {
						cnt++
											} else {
						res[i]+=cnt
											cnt=0
					}
				}
			}
		}
		return res
	}
	


func main() {
	s := "|**|*|****||"
	queries := [][]int{
		{1, 7},
		{0, 11},
		{1, 1},
	}
	fmt.Println("s=", s, "queries =", queries)
	fmt.Println(platesBetweenCandles(s, queries))
}
