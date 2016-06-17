/*
Soultion for Project Euler Problem #2 - https://projecteuler.net/problem=2

(c) 2016 dpetker
*/
package main

import "fmt"

func main() {
	step1 := 1
	step2 := 1
	sum := 0

	for step2 < 4000000 {
		newStep := step1 + step2

		if newStep%2 == 0 {
			sum += newStep
		}

		step1 = step2
		step2 = newStep
	}

	fmt.Println("Fibonacci sum:", sum)
}
