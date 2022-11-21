package main

import "fmt"

func isUgly(n int) bool {
	if n<1 {
		return false
	}
	x:=n
for i:=1; i<n; i++ {
if x%2==0{
	x=x/2
 } else if x%3==0{
	x=x/3
 }else if x%5==0{
	x=x/5
 } else {
 return false
 }
 if x==1 {
	return true
 }
}
return true
}

func main() {
	n := 7

	x := isUgly((n))
	fmt.Printf("For n=%v answer is %v", n, x)
}