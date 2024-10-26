package main

import (
	"os"

	"github.com/01-edu/z01"
)

// isPrime checks if a number n is prime.
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isSum calculates the sum of all prime numbers up to n.
func isSum(n int) int {
	sum := 0
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}

// printDigits recursively prints each digit of a number.
func printDigits(n int) {
	if n == 0 {
		return
	}
	printDigits(n / 10)             // Recursive call to process higher digits
	z01.PrintRune(rune(n%10 + '0')) // Print the last digit
}

// printNumber handles printing of the number.
func printNumber(n int) {
	if n == 0 {
		z01.PrintRune('0') // Special case for zero
		return
	}
	printDigits(n) // Call printDigits for non-zero numbers
}

// main function processes command-line arguments.
func main() {
	args := len(os.Args)

	if args != 2 {
		z01.PrintRune('0') // Print '0' for invalid argument count
		return
	}

	n := 0
	for _, char := range os.Args[1] {
		if char < '0' || char > '9' { // Check for valid digit characters
			z01.PrintRune('0') // Print '0' for invalid input
			return
		}
		n = n*10 + int(char-'0') // Convert char to int
	}

	// Note: The check for n < 0 is unnecessary since n cannot be negative.
	result := isSum(n)  // Calculate the sum of primes
	printNumber(result) // Print the result
	z01.PrintRune('\n') // Print a newline character
}
