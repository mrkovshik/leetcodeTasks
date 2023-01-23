package main

import "fmt"

func rotate(matrix [][]int) {
	l:=len(matrix)
for row:=0; row<l/2; row ++ {
	for col:= 0; col<l/2; col ++ {
matrix[row][col],matrix[row][l-1-col]=matrix[row][l-1-col],matrix[row][col]
fmt.Println(matrix)
matrix[row][col],matrix[l-1-row][col]=matrix[l-1-row][col],matrix[row][col]
fmt.Println(matrix)
matrix[l-1-row][l-1-col],matrix[l-1-row][col]=matrix[l-1-row][col],matrix[l-1-row][l-1-col]
fmt.Println(matrix)
	}
	}
}

func main() {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	for i:=range matrix {
		fmt.Printf("%v\n", matrix[i])
	}
	rotate(matrix)
	for i:=range matrix {
		fmt.Printf("%v\n", matrix[i])
	}
}