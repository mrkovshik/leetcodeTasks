package main

import "fmt"

func nearestExit(maze [][]byte, entrance []int) int {
crdx:=entrance[1]
crdy:=entrance[0]
}

func main() {
	maze := [][]byte{
		{43, 43, 46, 43},
		{46, 46, 46, 43},
		{43, 43, 43, 46},
	}
	entrance := []int{1, 2}

	res := nearestExit(maze, entrance)

	fmt.Printf("Maze is:\n%v\n%v\n%v\n Enterance at %v\n Answer is %v", maze[0], maze[1],maze[2], entrance, res)
}