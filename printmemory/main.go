package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func hexDegit2Rune(digit byte) rune {
	if digit < 10 {
		return rune(digit + '0')
	} else {
		return rune(digit - 10 + 'a')
	}
}

func PrintMemory(arr [10]byte) {
	for i, c := range arr {
		digit1 := c / 16
		digit2 := c % 16

		z01.PrintRune(hexDegit2Rune(digit1))
		z01.PrintRune(hexDegit2Rune(digit2))

		z01.PrintRune(' ')

		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}

	for _, char := range arr {
		if char >= 32 && char <= 126 {
			z01.PrintRune(rune(char))
		} else {
			z01.PrintRune('.')
		}

	}
}
