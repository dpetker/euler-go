/*
Soultion for Project Euler Problem #16 - https://projecteuler.net/problem=16

(c) 2016 dpetker
*/
package main

import (
	"fmt"
	"math/big"
)

func main() {
	var power = new(big.Int)
	power.Exp(big.NewInt(2), big.NewInt(1000), nil)

	sum := 0
	// range gives us runes, FYI
	for _, r := range power.String() {
		sum += int(r - '0')
	}

	fmt.Println(sum)
}
