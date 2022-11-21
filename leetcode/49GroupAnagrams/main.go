package main

import (
	"fmt"
	"strconv"
	)

func groupAnagrams(strs []string) [][]string {
res:=make([][]string,1)
var checksum,checkmult,k int
check:=make(map [string]int)
for i:=range strs{
	checkmult=1
checksum=0
for j:=0; j<len(strs[i]); j++{
	checksum+=int(strs[i][j])
	checkmult*=int(strs[i][j])
}
key:=strconv.Itoa(checkmult)+strconv.Itoa(checksum)
fmt.Println("checking word",strs[i])
fmt.Println("checkmult=",checkmult, "chchecksum=",checksum, "key=", key)
_,ok:=check[key]
if ok==true{
	fmt.Println("ok=",ok)
	res[check[key]]=append(res[check[key]], strs[i])
} else {
	fmt.Println("ok=",ok)
	check[key]=k
	
	if k==0{
		res[0]=append(res[0], strs[i])
	} else {
	v:=[]string{strs[i]}
	res=append(res,v )
	
	}
	k++
}
fmt.Println("map=",check)
}
return res
}

func main() {
	s := []string{"fsd", "dsf", "dfsg", "dfs"}
	fmt.Println("s=",s)
	fmt.Println(groupAnagrams(s))
}