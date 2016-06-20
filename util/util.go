/*
Utility functions for Project Euler

(c) 2016 dpetker
*/
package util

// Naive algorithm from: https://en.wikipedia.org/wiki/Primality_test#Pseudocode
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	} else if n <= 3 {
		return true
	} else if n%2 == 0 || n%3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}

	return true
}

func NumericByteToInt(b byte) int {
	return int(b - 48)
}
