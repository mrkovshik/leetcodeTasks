package main

import "fmt"
var res string

func convert(s string, numRows int) string { 
	if numRows ==1{
		return s
	}
add:=1
rowsSlice:= make([][]byte, numRows)
switcher:=0
for i:=0; i<len(s); i++ {
rowsSlice[switcher]=append(rowsSlice[switcher], s[i])
switcher=switcher+add
if switcher==numRows{
	switcher=switcher-2
	add=-1
}
if switcher==0{
	add=1
}
}
res:= ""
for j:=0; j<numRows; j++{
	res= res+string(rowsSlice[j])
	fmt.Println(" res= ", res)
}

return res
}

func main() {
	s := "abcdefg"
	n := 3
	newString := convert(s, n)
	fmt.Printf("source string = %v /n number of lines = %v /n converted string = %v", s,n,newString)
}