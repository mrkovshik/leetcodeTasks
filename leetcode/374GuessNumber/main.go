package main

import "fmt"

//   Forward declaration of guess API.
//   @param  num   your guess
//   @return 	     -1 if num is higher than the picked number
//  			      1 if num is lower than the picked number
//                otherwise return 0
func guess(num int) int {
	pick := 2
	switch {
	case num > pick:
		return -1
	case num < pick:
		return 1
	default:
		return 0
	}
}

func guessNumber(n int) int {
	base := n
	reply := 1
	if n != 1 {
		base = n / 2
	}
	trial := base

	for round := 1; round < n+2; round++ {
		reply = guess(trial)
		fmt.Printf("round=%v, base =%v, trial=%v, reply=%v", round, base, trial, reply)
		fmt.Println()
		if base == 1 {
			base = 2
		} else {
			base = base / 2
		}
		switch {
		case reply == 1:
			trial = trial + base
		case reply == -1:
			trial = trial - base
		default:
			return trial
		}
	}
	return trial
}

func main() {
	x := guessNumber(2)
	fmt.Printf("Pick =%v", x)
}
