package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args);i++{
sep =" "
s+=sep+os.Args[i]
    }  
    fmt.Println("s=",s)
	}
