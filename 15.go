/*
Soultion for Project Euler Problem #15 - https://projecteuler.net/problem=15

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

func main() {
	// Basically, we're going to create a 20x20 grid and fill it with Pascal's
	// Triangle :)
	GRID := make([][]int, 20)

	for i := 0; i < len(GRID); i++ {

		GRID[i] = make([]int, 20)
		for j := 0; j < len(GRID[i]); j++ {
			idx := i - 1
			left := 1
			if idx >= 0 {
				left = GRID[i-1][j]
			}

			idx = j - 1
			top := 1
			if idx >= 0 {
				top = GRID[i][j-1]
			}

			GRID[i][j] = left + top
		}

	}

	fmt.Println(GRID[19][19])
}
