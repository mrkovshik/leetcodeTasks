package main

import "fmt"
type quadMap map[byte]bool
	type colMap map[byte]bool


func defineQuad (row int, col int, quadsList [] quadMap) int {
	var quadNum int
	switch {
	case row < 3:
		switch {
		case col < 3:
			quadNum = 0
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 3 && col < 6:
			quadNum = 1
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 6 && col < 9:
			quadNum = 2
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		}

	case row >= 3 && row < 6:
		switch {
		case col < 3:
			quadNum = 3
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 3 && col < 6:
			quadNum = 4
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 6 && col < 9:
			quadNum = 5
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		}
	case row >= 6 && row < 9:
		switch {
		case col < 3:
			quadNum = 6
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 3 && col < 6:
			quadNum = 7
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		case col >= 6 && col < 9:
			quadNum = 8
			if quadsList[quadNum] == nil {
				quadsList[quadNum] = make(map[byte]bool, 9)
			}
		}
	}
	return quadNum
}

func isValidSudoku(board [][]byte) bool {
	
	var rowMap map[byte]bool
	quadsList := make([]quadMap, 9)
	colsList := make([]colMap, 9)

	var quadNum int
	for row := 0; row < 9; row++ {
		rowMap = make(map[byte]bool, 9)
				for col := 0; col < 9; col++ {
			quadNum=defineQuad(row,col,quadsList)
			if colsList[col]==nil {
				colsList[col]=make(map[byte]bool, 9)
			}
				x := board[row][col]
				fmt.Println("row = ", row, " col = ",col, "quad = ", quadNum, "x = ",x)
			if x != 0 {
				_, okQuad := quadsList[quadNum][x]
				if okQuad != true {
					quadsList[quadNum][x] = true
				} else {
					return false
				}
				_, okRow := rowMap[x]
				if okRow != true {
					rowMap[x] = true
				} else {
					return false
				}
				_, okCol := colsList[col][x]
				if okCol != true {
					colsList[col][x] = true
				} else {
					return false
				}

			}
			fmt.Println("rowMap",row,"=", rowMap[x], " colMap =",colsList[col][x], "quadMap",quadNum,"=", quadsList[quadNum], "x= ", x)
		}
	}
	return true
}

func main() {
	board := [][]byte{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 7},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	fmt.Println(board)
	fmt.Println(isValidSudoku(board))

}
