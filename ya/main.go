package main

import (
	"fmt"
)

func Mul(a interface{}, b int) interface{} {
var	res  interface{}
	switch a2 := a.(type) {
	case string: 
	res=a.(string)
	for i:=1; i<6; i++{
res=res.(string)+a.(string)
	}
case int: 
a2=a.(int)
res=a2*b
	}
return res
}


func main(){
a:=7
b:=3
	fmt.Println(Mul(a,b))
 }