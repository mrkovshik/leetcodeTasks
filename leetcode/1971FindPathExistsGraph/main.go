package main

import "fmt"

func validPath(n int, edges [][]int, source int, destination int) bool {
	var newSource int
		if source == destination{
		return true
		}
			fmt.Println("edges = ", edges)
	fmt.Println("source = ", source)
		for i:=0; i<len(edges);i++ {		
			fmt.Printf("searching %v in %v\n",source, edges[i])
		
			if edges[i][0]==source||edges[i][1]==source{
		if edges[i][0]==destination||edges[i][1]==destination{
			fmt.Printf("found %v in %v\n",destination, edges[i])
			return true
		} else {
		
			if edges[i][0]==source{
				newSource=edges[i][1]
				} else {
				newSource=edges[i][0]
			}	
			newEdges:=make([][]int,len(edges))
			copy(newEdges, edges)
			newEdges=append(newEdges[:i],newEdges[i+1:]...)
			if !validPath (n,newEdges,newSource,destination){
				continue
			} else {
				return true
			}
		}
			}
		}
		return false
	}


func main() {
	n := 10
	edges := [][]int{{4, 3}, {1, 4}, {4, 8},{1, 7},{6,4},{4,2},{7,4},{4,0},{0,9},{5,4}}
	
	source := 5
	destination := 9
	fmt.Printf(" n = %v,\n edges = %v\n from %v to %v \n answer is %v", n, edges,source,destination,validPath(n,edges,source,destination))

}