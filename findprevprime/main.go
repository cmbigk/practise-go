package main

import "fmt"

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func FindPrevPrime(nb int) int {
	for nb >= 2 {
		if isPrime(nb) {
			return nb
		}
		nb--
	}
	return 0
}

func main() {
	args := append([]int{99999}, 5, 4, 1, 0, 9, 8, 2, 1, 3, 4)
	for _, arg := range args {
		fmt.Println(FindPrevPrime(arg))
	}
}
