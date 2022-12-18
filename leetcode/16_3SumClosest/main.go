package main

import (
	"fmt"
	"math"
	)

func threeSumClosest(nums []int, target int) int {
	var min,res int
	var firstTime bool
numMap:=make (map [int] int)

for _,i:= range nums {
	numMap[i]++
	fmt.Println("map= ", numMap)
}
for i:= range numMap{
	fmt.Println("i= ", i)
	for j:= range numMap{
		fmt.Println("   j= ", j)
		if i==j&& numMap[i]==1 {
			continue
		}
		for k:= range numMap{
			fmt.Println("    k= ", k)
			if (k==i&&numMap[k]==1)||(k==j&&numMap[k]==1)||(k==j&&k==i&&numMap[k]==2){
				continue
			}
			dif:=int(math.Abs(float64(target-(i+j+k))))
		
			if !firstTime{
				min=dif
				fmt.Println("min= ", min)
				fmt.Println("dif= ", dif)
				res=i+j+k
				fmt.Println("res= ", res)

				firstTime=true
			}
			if dif<min{
				min=dif
				fmt.Println("min= ", min)
				fmt.Println("dif= ", dif)
				res=i+j+k
				fmt.Println("res= ", res)
			}
		
		}
	}
}
return res
}

func main() {
	nums := []int{1, 0, 5, 2, 0, 0, 0, 0, 1, -1}
	target := -3
	fmt.Println(threeSumClosest(nums, target))
}