/*
Soultion for Project Euler Problem #17 - https://projecteuler.net/problem=17

(c) 2016 dpetker
*/
package main

import (
	"fmt"
)

var baseDict = map[int]int{
	1:    len("one"),
	2:    len("two"),
	3:    len("three"),
	4:    len("four"),
	5:    len("five"),
	6:    len("six"),
	7:    len("seven"),
	8:    len("eight"),
	9:    len("nine"),
	10:   len("ten"),
	11:   len("eleven"),
	12:   len("twelve"),
	13:   len("thirteen"),
	14:   len("fourteen"),
	15:   len("fifteen"),
	16:   len("sixteen"),
	17:   len("seventeen"),
	18:   len("eighteen"),
	19:   len("nineteen"),
	20:   len("twenty"),
	30:   len("thirty"),
	40:   len("fourty"),
	50:   len("fifty"),
	60:   len("sixty"),
	70:   len("seventy"),
	80:   len("eighty"),
	90:   len("ninety"),
	100:  len("hundred"),
	1000: len("thousand"),
}

func main() {
	sum := 0

	for i := 1; i < 1001; i++ {
		sum += findCharCount(i)
	}

	fmt.Println(sum)
}

func findCharCount(num int) int {
	if num <= 20 {
		return baseDict[num]
	}

	return 0
}
