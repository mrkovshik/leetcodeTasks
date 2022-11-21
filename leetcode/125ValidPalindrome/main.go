package main

import (
	"fmt"
	"strings"
)
func checkI (i *int, s string) {
if (s[*i]<97&&s[*i]>57)||s[*i]>122||s[*i]<48 {
*i++
if *i<len(s)-1{
checkI(i,s)
}
}
}
func checkJ (j *int, s string) {
	if (s[*j]<97&&s[*j]>57)||s[*j]>122||s[*j]<48 {
	*j--
	if *j>0 {
	checkJ(j,s)
	
	}
}
}
func isPalindrome(s string) bool {
	s=strings.ToLower(s)
	if len(s)==0 || len(s)==1 {
		return true
	}
j:=len(s)-1
for i:=0; i<=j; i++ {
	checkI(&i, s)
	checkJ(&j, s)
fmt.Println("i=", i, "j=", j)
if i==len(s)-1&&j==0{
	return true
}
if s[i]!=s[j]  {
	return false
}
if j>0{
j--
}
}
return true
}

func main() {
	s := "\"Stop!\" nine myriad murmur. \"Put up rum, rum, dairymen, in pots.\""
	fmt.Println(isPalindrome(s))
}