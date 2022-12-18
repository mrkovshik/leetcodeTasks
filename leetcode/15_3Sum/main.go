package main

import "fmt"

func threeSum(nums []int) [][]int {
	numMap:= make(map[int]int, 0)
	res:= make([][]int, 0)

for _,i:= range nums {
	numMap[i]++
}

for i:= range numMap{
	
	for j:= range numMap{
		
		if i>j || (i==j&&numMap[i]==1){
			continue
		}
		k:=-(i+j)
		if k<i || k<j {
			continue
		}
		if n,ok:=numMap[k]; ok {
			if (i==k && n ==1)|| (j==k&& n==1)||(i==k&&j==k&& n==2){
				continue
			}
			res= append(res, []int{i,j,k})
		
		}
	}
}
return res
}

func main() {
	nums := []int{1, 0, -1, 0, 0, 0,0,0,1,-1}
	fmt.Println(threeSum(nums))
}