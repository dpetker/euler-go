/*
Soultion for Project Euler Problem #7 - https://projecteuler.net/problem=7

(c) 2016 dpetker
*/
package main

import (
	"euler-go/util"
	"fmt"
)

func main() {
	currPrime := 6

	for i := 15; ; i += 2 {
		if util.IsPrime(i) {
			currPrime++

			if currPrime == 10001 {
				fmt.Println(i)
				return
			}
		}
	}
}
