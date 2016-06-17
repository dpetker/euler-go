/*
Soultion for Project Euler Problem #1 - https://projecteuler.net/problem=1

(c) 2016 dpetker
*/
package main

import "fmt"

func main() {
  multiples := make([]int, 0, 200)

  for i := 0; i < 1000; i++ {
    if (i % 3 == 0) || (i % 5 == 0) {
      multiples = append(multiples, i)
    }
  }

  sum := 0
  for _, value := range multiples {
    sum += value
  }

  fmt.Println("Sum of multiples:", sum)
}
