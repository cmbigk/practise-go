package main

import "fmt"

func Gcd(a, b uint) uint {
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))

	fmt.Println(Gcd(12, 23))
	fmt.Println(Gcd(25, 15))
	fmt.Println(Gcd(23043, 122))
	fmt.Println(Gcd(11, 77))
}

/*
package main

import "fmt"

// Gcd function calculates the greatest common divisor of two uint numbers.
func Gcd(a, b uint) uint {
	// If any number is 0, return 0 as the result.
	if a == 0 || b == 0 {
		return 0
	}

	// Use Euclid's algorithm to compute GCD.
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Test cases
	fmt.Println(Gcd(42, 10))  // Output: 2
	fmt.Println(Gcd(42, 12))  // Output: 6
	fmt.Println(Gcd(14, 77))  // Output: 7
	fmt.Println(Gcd(17, 3))   // Output: 1

	fmt.Println(Gcd(12, 23))  // Output: 1
	fmt.Println(Gcd(25, 15))  // Output: 5
	fmt.Println(Gcd(23043, 122)) // Output: 1
	fmt.Println(Gcd(11, 77))  // Output: 11
}
*/
