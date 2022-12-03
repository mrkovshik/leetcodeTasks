package main

import "fmt"

func maxArea(height []int) int {
qnty:=len(height)
var area, max int
for i:=0; i<qnty-1; i++{
	for j:=i+1; j<qnty;j++{
if height[i]>height[j]{
	area=height[j]*(j-i)
} else {
	area=height[i]*(j-i)
}
if area>max{
	max=area
}
	}
}	

return max
}

func main() {
	height := []int{3, 4, 2, 5}
	max := maxArea(height)
	fmt.Printf("height= %v max Area =%v", height, max)
}