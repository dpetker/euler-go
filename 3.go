/*
Soultion for Project Euler Problem #3 - https://projecteuler.net/problem=3

Algorithm inspiration from http://stackoverflow.com/questions/23287/largest-prime-factor-of-a-number

(c) 2016 dpetker
*/
package main

import "fmt"

func main() {
	fmt.Println(findMaxFactor(600851475143))
}

func findMaxFactor(testVal int) int {
	divisor := 2
	currMax := 0

	for testVal > 1 {
		for testVal%divisor == 0 {

			if divisor > currMax {
				currMax = divisor
			}

			testVal /= divisor
		}

		divisor++

		if divisor*divisor > testVal {
			if testVal > 1 && testVal > currMax {
				currMax = testVal
			}
			return currMax
		}
	}

	return currMax
}
