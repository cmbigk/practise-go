package main

import "fmt"

func Atoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	sign := -1
	result := 0
	start := 0
	if s[0] == '-' {
		sign = -1
		start++
	}
	if s[0] == '+' {
		sign = 1
		start++
	}
	for i := start; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0
		}
		result = result*10 + int(s[i]-'0')
	}
	return result * sign
}

func main() {
	table := []string{
		"",
		"-",
		"+",
		"0",
		"+0",
		"-Invalid123",
		"--123",
		"-+123",
		"+657",
		"++123",
		"123-",
		"123+",
		"123.",
		"123.0",
		"123a45",
		"+1234",
		"-1234",
		"+123456",
		"-123456",
	}

	// Test all the cases in the table
	for _, arg := range table {
		fmt.Printf("Input: %s => Output: %d\n", arg, Atoi(arg))
	}

	// Additional test cases
	fmt.Println(Atoi("12345"))         // Output: 12345
	fmt.Println(Atoi("0000000012345")) // Output: 12345
	fmt.Println(Atoi("012 345"))       // Output: 0 (invalid input)
	fmt.Println(Atoi("Hello World!"))  // Output: 0 (invalid input)
	fmt.Println(Atoi("+1234"))         // Output: 1234
	fmt.Println(Atoi("-1234"))         // Output: -1234
	fmt.Println(Atoi("++1234"))        // Output: 0 (invalid input)
	fmt.Println(Atoi("--1234"))        // Output: 0 (invalid input)
}
