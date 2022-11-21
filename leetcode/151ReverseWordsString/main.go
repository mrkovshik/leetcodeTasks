package main

import "fmt"

func reverseWords(s string) string {
var res,word string
start:=false
for i:=0; i<len(s);i++ {
	switch{
	case s[i] != 32 && start==false:
		 word=string(s[i])
		 start=true
	case s[i]!=32 && start==true:
		word=word + string(s[i])
	case s[i]==32 && start==true: 
	if res!="" {
	res=word+" "+res
	} else {
		res=word
	}
	start=false
	word=""	
case s[i] == 32 && start==false:
	continue
}

fmt.Println("word=", word)
	fmt.Println("res=", res)
}
if word!=""&&res!=""{
res=word+" "+res
}
if word!=""&&res==""{
	res=word
	}
return res
}

func main() {
	s := "narkoman"
	fmt.Println("s=",s)
	fmt.Println("res=", reverseWords(s))
}