package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

for i:=2; i<len(tokens); i++ {
	fmt.Printf("tokens = %v\n", tokens)
	fmt.Printf("i = %v  t = %v \n", i, tokens[i])
	if tokens[i]=="+"||tokens[i]=="-"||tokens[i]=="*"||tokens[i]=="/"{
		x,_:=strconv.Atoi(tokens[i-2])
	y,_:=strconv.Atoi(tokens[i-1])
	switch tokens[i]{
	case "+" : 
	tokens[i]= strconv.Itoa (x+y)
	case "-" : 
	tokens[i]= strconv.Itoa (x-y)
	case "*" : 
	tokens[i]= 	strconv.Itoa (x*y)
		case "/" : 
	tokens[i]= strconv.Itoa (x/y)
		}
	tokens = append(tokens[:i-2], tokens[i:]... )
	i=i-2
}

	}
res,_:=strconv.Atoi(tokens[0])
return res
}



func main() {

	tokens := []string{"2","1","+","3","*"}
	// fmt.Printf("tokens = %v\n", tokens)
	fmt.Printf("answer = %v", evalRPN(tokens))
}