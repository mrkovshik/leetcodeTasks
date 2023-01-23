package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	var x string = "11"
	var counter int = 1
if n == 1 {
	return "1"
}
for i:=2; i<n; i++{
	fmt.Println("x=" , x)
numSlice:= strings.Split(x,"")
x=""
for j:=1;j<len(numSlice);j++ {
	if numSlice[j]==numSlice[j-1]{
	counter++
	if j!=len(numSlice)-1{
continue
	}
	} 
	
x=x+strconv.Itoa(counter)+numSlice[j-1]
counter=1
if j==len(numSlice)-1 && numSlice[j]!=numSlice[j-1]{

	x=x+"1"+numSlice[j]
	
}
}
}
return x
}


func main() {
	n := 7
	
	fmt.Printf("for n =%v answer is %v", n, countAndSay(n))
}