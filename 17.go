/*
Soultion for Project Euler Problem #17 - https://projecteuler.net/problem=17

(c) 2016 dpetker
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var baseDict = map[int]int{
	0:    0,
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
	40:   len("forty"),
	50:   len("fifty"),
	60:   len("sixty"),
	70:   len("seventy"),
	80:   len("eighty"),
	90:   len("ninety"),
	100:  len("hundred"),
	1000: len("onethousand"),
}

func main() {
	var sum uint32
	var wg sync.WaitGroup

	wg.Add(1000)

	for i := 1; i < 1001; i++ {
		go parallelWordCounter(i, &sum, &wg)
	}

	wg.Wait()

	fmt.Println(sum)
}

func parallelWordCounter(num int, accum *uint32, wg *sync.WaitGroup) {
	atomic.AddUint32(accum, uint32(findCharCount(num)))
	wg.Done()
}

func findCharCount(num int) int {
	switch {
	case num <= 20 || num == 1000:
		return baseDict[num]
	case num > 20 && num < 100:
		tens := (num / 10) * 10
		ones := num % 10
		return baseDict[tens] + baseDict[ones]
	case num >= 100 && num < 1000:
		hundreds := num / 100
		remainder := num % 100

		// +3 letters for "and" as in, "one hundred and forty two"
		tempSum := baseDict[hundreds] + baseDict[100] + 3 + findCharCount(remainder)

		if remainder == 0 {
			// No "and" when we're right on a hundred
			tempSum -= 3
		}

		return tempSum
	}

	return 0
}
