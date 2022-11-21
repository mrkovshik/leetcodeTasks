package main

import "fmt"

//  * Forward declaration of isBadVersion API.
//  * @param   version   your guess about first bad version
//  * @return 	 	      true if current version is bad
//  *			          false if current version is good

func isBadVersion(version int) bool {
	BadVersion := 3
	if version >= BadVersion {
		return true
	}
	return false
}

func firstBadVersion(n int) int {
	guess := n / 2
	pitch := n / 4
	k := 0
	for i := 0; k == 0; i++ {
		fmt.Printf("round =%v,  guess=%v, reply=%v", i,guess,isBadVersion(guess))
		fmt.Println()
		if isBadVersion(guess) {
			guess -= pitch
		} else {
			guess += pitch
		}

		pitch = pitch / 2
		if pitch < 1 {
			x:=isBadVersion(guess)
			y:=isBadVersion(guess)
			for j:=1; x==y; j++{
			if isBadVersion(guess) {
				guess--
				x=isBadVersion(guess)
			} else {
				guess++
				x=isBadVersion(guess)
			}
		}
		if !x{
			guess++
		}
			return guess
		}
	}
	return guess
}
func main() {
	n:= 4
	x := firstBadVersion(n)
	fmt.Printf("n=%v, first bad=%v", n, x)

}
