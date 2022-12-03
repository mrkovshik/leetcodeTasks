package main

import (
		"fmt"
	"strconv"
)

func reverse(x int) int {
strInt:=strconv.Itoa(x)
lng:=len(strInt)
newStr:=make([] byte, 0)
zeroStart:=true
for i:=lng-1; i>=0; i-- {
	fmt.Println("i= ",i)
	fmt.Println("strInt[i]= ",strInt[i])
if strInt[i]==48 && zeroStart==true {
continue
} else {
	zeroStart=false
}
newStr=append (newStr, strInt[i])
fmt.Println(newStr)
}
res,_:=strconv.Atoi(string(newStr))

return res
}

func main() {
	x := 123
	res := reverse(x)
	fmt.Printf("x =%v, reverse =%v", x,res)
}