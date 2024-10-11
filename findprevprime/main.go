package main

import "fmt"

func isPrime(nb int) bool {
	if nb < 2 {
		return false // Add a return statement for the case when nb < 2
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
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
	args := append([]int{99999}, 5, 4, 1, 0, 9, 18, 62, 16, 83, 4)
	for _, arg := range args {
		fmt.Println(FindPrevPrime(arg))
	}
}
