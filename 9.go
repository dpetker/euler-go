/*
Soultion for Project Euler Problem #9 - https://projecteuler.net/problem=9

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

func main() {
	// List of triples from https://en.wikipedia.org/wiki/Pythagorean_triple
	TRIPLES := [][]int{
		[]int{3, 4, 5},
		[]int{5, 12, 13},
		[]int{8, 15, 17},
		[]int{7, 24, 25},
		[]int{20, 21, 29},
		[]int{12, 35, 37},
		[]int{9, 40, 41},
		[]int{28, 45, 53},
		[]int{11, 60, 61},
		[]int{16, 63, 65},
		[]int{33, 56, 65},
		[]int{48, 55, 73},
		[]int{13, 84, 85},
		[]int{36, 77, 85},
		[]int{39, 80, 89},
		[]int{65, 72, 97},
		[]int{20, 99, 101},
		[]int{60, 91, 109},
		[]int{15, 112, 113},
		[]int{44, 117, 125},
		[]int{88, 105, 137},
		[]int{17, 144, 145},
		[]int{24, 143, 145},
		[]int{51, 140, 149},
		[]int{85, 132, 157},
		[]int{119, 120, 169},
		[]int{52, 165, 173},
		[]int{19, 180, 181},
		[]int{57, 176, 185},
		[]int{104, 153, 185},
		[]int{95, 168, 193},
		[]int{28, 195, 197},
		[]int{84, 187, 205},
		[]int{133, 156, 205},
		[]int{21, 220, 221},
		[]int{140, 171, 221},
		[]int{60, 221, 229},
		[]int{105, 208, 233},
		[]int{120, 209, 241},
		[]int{32, 255, 257},
		[]int{23, 264, 265},
		[]int{96, 247, 265},
		[]int{69, 260, 269},
		[]int{115, 252, 277},
		[]int{160, 231, 281},
		[]int{161, 240, 289},
		[]int{68, 285, 293},
	}

	// Also from the above wikipedia article we know that scalar multiples of
	// the above triples are themselves triples. This helps confine our solution
	// space by quite a bit
	result := make(chan int)
	for _, triple := range TRIPLES {
		go checkTuple(triple[0], triple[1], triple[2], result)
	}

	product := <-result
	fmt.Println(product)
}

func checkTuple(a int, b int, c int, result chan int) {
	sum := a + b + c
	if sum == 1000 {
		result <- a * b * c
	} else if sum < 1000 {
		for i := 2; sum < 1000; i++ {
			sum = (a * i) + (b * i) + (c * i)
			if sum == 1000 {
				result <- (a * i) * (b * i) * (c * i)
			}
		}
	}
}
