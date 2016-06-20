/*
Soultion for Project Euler Problem #6 - https://projecteuler.net/problem=6

(c) 2016 dpetker
*/
package main

import "fmt"

// We don't *really* need goroutines for this, as there are constant-time
// formulae to calculate both numbers, but this was a good exercise for me :)
func main() {
	TEST_VAL := 100
	c1 := make(chan int)
	c2 := make(chan int)

	go findSumOfSquares(TEST_VAL, c1)
	go findSquareOfSum(TEST_VAL, c2)

	sum1, sum2 := <-c1, <-c2

	fmt.Println(sum2 - sum1)
}

func findSumOfSquares(n int, c chan int) {
	// Formula from: https://en.wikipedia.org/wiki/Square_pyramidal_number
	c <- (n * (n + 1) * ((2 * n) + 1)) / 6
}

func findSquareOfSum(n int, c chan int) {
	// Formula from: https://en.wikipedia.org/wiki/Triangular_number
	result := (n * (n + 1)) / 2
	c <- result * result
}
