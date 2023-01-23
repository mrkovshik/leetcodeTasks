package main

import (
	"fmt"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	var res int
	sortedSlice:=make ([]int,0)
	spaceMap:=make(map [int] int,len(rocks))
for i:=0;i<len(rocks);i++ {

	space:=capacity[i]-rocks[i]
	fmt.Printf("sortedSlice = %v , Map = %v \n", sortedSlice, spaceMap)
	fmt.Printf("cap = %v, rocks = %v, empty space = %v\n", capacity[i], rocks[i], space)
	
	if space==0{
		res++
		continue
	}
	if _,ok:= spaceMap[space]; !ok {
		spaceMap [space]=1
		switch{
		case len(sortedSlice)==0||space>=sortedSlice[len(sortedSlice)-1]:
			sortedSlice= append(sortedSlice, space)
		case space<sortedSlice[0]:
			sortedSlice = append([]int{space},sortedSlice... )
		
		default:	
		n:=len(sortedSlice)		
for j:=1; j<n; j++{
			if space <= sortedSlice[j]{
				sortedSlice = append(sortedSlice[:j+1],sortedSlice[j:]... )
sortedSlice[j]=space		
				break
		}
		}
	}
	} else {
		spaceMap [space]++
	}
	}
	fmt.Printf("sortedSlice = %v , Map = %v \n", sortedSlice, spaceMap)
	fmt.Println("res = ", res)
	for _,j:=range sortedSlice{
		if additionalRocks>= j*spaceMap[j]{
			additionalRocks-=j*spaceMap[j]
			res+=spaceMap[j]
			continue
		}else{
			for additionalRocks>=j{
				res++
				additionalRocks-=j
			}
			break
		
		
		}
}
return res
}

func main() {
	cap := []int{54,18,91,49,51,45,58,54,47,91,90,20,85,20,90,49,10,84,59,29,40,9,100,1,64,71,30,46,91}
	rocks := []int{14,13,16,44,8,20,51,15,46,76,51,20,77,13,14,35,6,34,34,13,3,8,1,1,61,5,2,15,18}
	add := 77
	fmt.Printf("You can fill %v bags", maximumBags(cap,rocks,add))
}