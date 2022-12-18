package main

import (
		"fmt"
	"strconv"
	"math"
)

func reverse(x int) int {
	
strInt:=strconv.Itoa(int(math.Abs(float64(x))))
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
if x<0 {
	res=-res
}
if res > 2147483647 || res<(-2147483648) {
	return 0
}

return res
}

func main() {
	x := 58
	res := reverse(x)
	fmt.Printf("x =%v, reverse =%v", x,res)
}