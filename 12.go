/*
Soultion for Project Euler Problem #12 - https://projecteuler.net/problem=12

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

func main() {
	numDivisors := 0
	test := 0

	for i := 1; numDivisors <= 501; i++ {
		test = generateTriangleNum(i)
		numDivisors = countDivisors(test)
	}

	fmt.Println(test)
}

func generateTriangleNum(n int) int {
	return (n * (n + 1)) / 2
}

func countDivisors(n int) int {
	limit := n
	count := 0

	if n == 1 {
		return 1
	}

	for i := 1; i < limit; i++ {
		if n%i == 0 {
			limit = n / i
			if limit != i {
				count++
			}

			count++
		}
	}

	return count
}
