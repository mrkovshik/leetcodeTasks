package main

import "fmt"

func isIsomorphic(s string, t string) bool {
	charMapS := make(map[byte]byte, 0)
	charMapT := make(map[byte]byte, 0)
	for i := range s {

		if _, okS := charMapS[s[i]]; !okS {
			charMapS[s[i]] = t[i]
		}
		if _, okT := charMapT[t[i]]; !okT {
			charMapT[t[i]] = s[i]
		}
		if charMapT[t[i]] != s[i] || charMapS[s[i]] != t[i] {
			return false
		}

	}
	return true
}

func main() {
	s := "egg"
	t := "add"

	fmt.Print(isIsomorphic(s, t))

}
