package main

import (
	"fmt"
	"math"
)

func countPrimes(n int) int {
	primes := []int{3}
	if n <= 2 {
		return 0
	}
	if n == 3 {
		return 1
	}
OuterLoop:
	for i := 5; i < n; i += 2 {
		sqrtN := int(math.Sqrt(float64(i))) + 1
		for j := 0; primes[j] < sqrtN; j++ {
			if i%primes[j] == 0 {
				continue OuterLoop
			}

		}
		primes = append(primes, i)
	}
	return len(primes) + 1
}

func main() {
	n := 10
	fmt.Print(countPrimes(n))
}
