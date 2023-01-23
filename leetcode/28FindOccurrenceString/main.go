package main

import (
	"fmt"
	
)

func strStr(haystack string, needle string) int {
	if len(needle)==0{
		return 0
	}
	k:=0
	for i:=range haystack{

if haystack[i]!=needle[k]{
	k=0
continue
} 
k++
if k==len(needle){
	return i-k+1
}
	}
	return -1
}

func main() {
	haystack := "k"
	needle := "g"
	fmt.Printf("You can find word %v at %v simbol of %v", needle, strStr(haystack,needle),haystack)
}