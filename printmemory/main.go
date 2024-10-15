package main

import (
	"github.com/01-edu/z01"
)

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func hexa2rune(n byte) rune {
	if n < 10 {
		return rune(n + '0')
	} else {
		return rune(n - 10 + 'a')
	}
}

func PrintMemory(arr [10]byte) {
	for i, c := range arr {
		firstdigit := c / 16
		seconddigit := c % 16

		z01.PrintRune(hexa2rune(firstdigit))
		z01.PrintRune(hexa2rune(seconddigit))

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
