package main

import "fmt"

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	result := ""
	for n > 0 {
		addingLastDigit := n % 10
		result = string(addingLastDigit+'0') + result
		n = n / 10
	}

	// Add the negative sign if needed
	if isNegative {
		result = string('-') + result
	}

	return result
}

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}
