package main

import (
	"fmt"
)

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	var cx1,cy1,cx2,cy2 int
if ax1<=bx1{
	if ax2<=bx1{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cx1=bx1
} else {
	if ax1>=bx2{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cx1=ax1
}
fmt.Println("cx1=", cx1)

if ax2>=bx2{
	if ax1>=bx2{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cx2=bx2
} else {
	if ax2<=bx1{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cx2=ax2
}

fmt.Println("cx2=", cx2)
if ay1<=by1{
	if ay2<=by1{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cy1=by1
} else {
	if ay1>=by2{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cy1=ay1
}
fmt.Println("cy1=", cy1)

if ay2>=by2{
	if ay1>=by2{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cy2=by2
} else {
	if ay2<=by1{
		return (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)
	} 
	cy2=ay2
}
fmt.Println("cy2=", cy2)

intrsArea:= (ax2-ax1)*(ay2-ay1)+(bx2-bx1)*(by2-by1)-(cx2-cx1)*(cy2-cy1)
return intrsArea
}

func main() {
	ax1 := -2
	ay1 := -2
	ax2 := 2
	ay2 := 2
	bx1 := -2
	by1 := -4
	bx2 := 2
	by2 := -2
	fmt.Printf("ax1=%v, ay1=%v, ax2=%v, ay2=%v, bx1=%v by1=%v, bx2=%v, by2=%v", ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Println()
	x:=computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Println("Area = ", x)
	
}