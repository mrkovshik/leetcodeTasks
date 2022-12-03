package main

import (
	"fmt"

)

func findMedianortedArrays(nums1 []int, nums2 []int) float64 {
	var newArr []int
	var j,k int
	var median float64

	for i:=0; i<len(nums1)+len(nums2); i++ {
		if j==len(nums1){
			newArr=append(newArr, nums2[k:]...)
			break
		}
		if k==len(nums2){
			newArr=append(newArr, nums1[j:]...)
			break
		}
		fmt.Println("i=",i)
if nums1[j]<=nums2[k]{
	newArr=append(newArr,nums1[j])
fmt.Println("newArr=",newArr)
j++
} else {
	newArr=append(newArr,nums2[k])
	fmt.Println("newArr=",newArr)
	k++
}

	}
	fmt.Println("newArr=",newArr)
	lng:=len(newArr)
if lng%2==0{
median= (float64(newArr[lng/2-1])+float64(newArr[lng/2]))/2
} else {
	median= float64(newArr[lng/2])
}


	return median
}
func main() {
nums1:=[]int{1,3}
nums2:=[]int{2}
fmt.Println("nums1=",nums1)
fmt.Println("nums2=",nums2)
x:=findMedianortedArrays(nums1,nums2)
fmt.Println("median =",x)
}