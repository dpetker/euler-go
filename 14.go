/*
Soultion for Project Euler Problem #14 - https://projecteuler.net/problem=14

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

func main() {
	maxLen := 0
	maxNum := 0

	for i := 1000000; i > 1; i-- {
		temp := computeChain(i)

		if temp > maxLen {
			maxLen = temp
			maxNum = i
		}
	}

	fmt.Println(maxNum)
}

func computeChain(n int) int {
	chain := 0

	for n > 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}

		chain++
	}

	return chain
}
