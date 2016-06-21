/*
Soultion for Project Euler Problem #10 - https://projecteuler.net/problem=10

(c) 2016 dpetker
*/
package main

import (
	"euler-go/util"
	"fmt"
)

func main() {
	sum := 0

	// This is *really* brute force, and we could easily parallelize it using
	// goroutines, but then we'd need to figure out when all 2 million checks
	// have completed. Given that the brute force solution runs in ~1s on my
	// laptop, I'm not inclined to over-optimize
	for i := 2; i < 2000000; i++ {
		if util.IsPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
