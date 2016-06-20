/*
Soultion for Project Euler Problem #4 - https://projecteuler.net/problem=4

(c) 2016 dpetker
*/
package main

import (
	"fmt"
	"strconv"
)

func main() {
	max := 0
	for i := 999; i > 0; i-- {
		for j := 999; j > 0; j-- {
			prod := i * j
			if prod > max && test(prod) {
				max = prod
			}
		}
	}

	fmt.Println(max)
}

func test(num int) bool {
	digits := []byte(strconv.Itoa(num))
	size := len(digits)

	for i := 0; i < size/2; i++ {
		if digits[i] != digits[size-1-i] {
			return false
		}
	}

	return true
}
