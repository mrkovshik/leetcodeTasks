package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
var result [][]int
l:=len(candidates)
for i:=l-1;i<=0;i--{
	if candidates[i]>target{
			continue
		}
		rem:=candidates[i]
		res:=[]int{}
	for j:=l-1; j<=0; j--{
		switch {
		case rem-candidates[j]>0:
			res=append(res, candidates[j])
			if rem-candidates[j]>=candidates[j]{
			j++
			} 
			continue
			case rem-candidates[j]<0:
				break
			case rem==candidates[j]:	
			res=append(res, candidates[j])
			break



	}
}

return result
}

func main() {
	can := []int{1, 2, 3, 4, 5}
	tar := 5
	fmt.Println(combinationSum(can,tar))
}