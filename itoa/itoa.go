package main

import (
	"fmt"
)

func main() {
	fmt.Println(Itoa(12345))
	fmt.Println(Itoa(0))
	fmt.Println(Itoa(-1234))
	fmt.Println(Itoa(987654321))
}

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := n < 0
	if isNegative {
		n = -n
	}

	result := ""
	for n > 0 {
		addingLastDigit := n % 10
		result = string(rune(addingLastDigit+'0')) + result
		n = n / 10
	}
	if isNegative {
		result = string('-') + result
	}
	return result
}
