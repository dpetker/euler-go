/*
Soultion for Project Euler Problem #18 - https://projecteuler.net/problem=18

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

func main() {
	FIELD := [][]int{
		[]int{8, 5, 9, 3},
		[]int{2, 4, 6, 0},
		[]int{7, 4, 0, 0},
		[]int{3, 0, 0, 0},
	}

	maxPaths := [][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	}

	for i := 0; i < len(FIELD); i++ {
		for j := 0; j < len(FIELD[i]); j++ {
			maxPaths[i][j] = findMaxPath(FIELD, i, j)
		}
	}

	fmt.Println(maxPaths[len(FIELD)-1][0])
}

func findMaxPath(field [][]int, row int, col int) int {
	return field[row][col]
}
