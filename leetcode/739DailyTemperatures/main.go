package main

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	var res []int
	var big struct{
		val int
		less int
		index int
	}
for i:=0; i<len(temperatures)-1; i++{
	fmt.Printf("i = %v \n", i)
	if temperatures[i]<big.val&& temperatures[i]>= big.less&& i<big.index {
		
		res=append(res, big.index-i)
		fmt.Printf("res(b) = %v\n", res)
		continue
	}
	
	
for j:=i+1;j<len(temperatures);j++ {
	fmt.Printf("j = %v\n", j)
if temperatures[j]>temperatures[i]{
	big.val,big.index,big.less=temperatures[j],j,temperatures[i]
	fmt.Printf("big = %v\n", big)
	res=append(res, j-i)
	fmt.Printf("res = %v\n", res)
	break
}
if j==len(temperatures)-1{
res=append(res, 0)
fmt.Printf("res = %v\n", res)
}
}
}
res=append(res, 0)
fmt.Printf("res = %v\n", res)
return res
}

func main() {
	temps := []int{34,80,80,34,34,80,80,80,80,34}
	fmt.Printf("temperatures = %v \n answer = %v\n", temps, dailyTemperatures(temps))
}