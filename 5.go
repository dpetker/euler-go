/*
Soultion for Project Euler Problem #5 - https://projecteuler.net/problem=5

(c) 2016 dpetker
*/
package main

import "fmt"

func main() {
	for i := 2520; ; i += 20 {
		if test(i) {
			fmt.Println(i)
			return
		}
	}
}

func test(num int) bool {
	for i := 1; i <= 20; i++ {
		if num%i != 0 {
			return false
		}
	}

	return true
}
