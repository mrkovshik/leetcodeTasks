package main

import "fmt"

func possibleBipartition(n int, dislikes [][]int) bool {
	if n ==1 || len(dislikes)==0{
			return true
	}
group1:=make(map [int] bool)
group2:=make(map [int] bool)
	group1[dislikes[0][0]]=false
	group2[dislikes[0][1]]=false
	fmt.Printf("group 1 = %v \ngroup 2 = %v\n", group1, group2)
for i:= 1; i<len(dislikes); i++ {
	
	pair:=dislikes[i]
_,ok0_1:= group1[pair[0]]	
_,ok0_2:= group2[pair[0]]
_,ok1_2:= group2[pair[1]]
_,ok1_1:= group1[pair[1]]
switch{
case !ok0_2&&ok0_1: 
// if ok1_2{
// 	return false
// }
group2[pair[1]]=false
case ok0_2&&!ok0_1:
	// if  ok1_1{
	// 	return false
	// }
	group1[pair[1]]=false
case !ok0_1&&!ok0_2:
	_,ok1_1:= group1[pair[1]]
	_,ok1_2:= group2[pair[1]]
	switch{
	case ok1_1 :
group2[pair[0]]=false
	case ok1_2:
group2[pair[1]]=false
group1[pair[0]]=false
	case !ok1_1&& !ok1_2:
if i<len(dislikes)-1{
	dislikes[i],dislikes[i+1]=dislikes[i+1],dislikes[i]
	i--
	continue
} 

}
}
if ok0_1&&ok1_1||ok0_2&&ok1_2{
	return false
}
fmt.Printf("group 1 = %v \ngroup 2 = %v\n", group1, group2)
}
	return true
}

func main() {
	n := 4
	dislikes := [][]int{{1, 2}, {1, 3}, {2, 3}}

	fmt.Printf("answer is %v", possibleBipartition(n,dislikes))
}