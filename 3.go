/*
Soultion for Project Euler Problem #3 - https://projecteuler.net/problem=3

Algorithm inspiration from http://stackoverflow.com/questions/23287/largest-prime-factor-of-a-number

(c) 2016 dpetker
*/
package main

import "fmt"

func main() {
	factors := findFactors(600851475143)

	fmt.Println(factors)
}

func findFactors(testVal int) []int {
	divisor := 2
	factors := make([]int, 0, 1000)

	for testVal > 1 {
		for testVal%divisor == 0 {
			factors = append(factors, divisor)
			testVal /= divisor
		}

		divisor++

		if divisor*divisor > testVal {
			if testVal > 1 {
				factors = append(factors, testVal)
			}
			return factors
		}
	}

	return factors
}
