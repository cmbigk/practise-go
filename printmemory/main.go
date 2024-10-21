package main

import (
	"github.com/01-edu/z01"
)

func main() {

	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func hex2Rune(s byte) rune {
	if s < 10 {
		return rune(s + '0')
	} else {
		return rune(s - 10 + 'a')
	}
}

func PrintMemory(arr [10]byte) {
	for i, c := range arr {
		firstdigit := c / 16
		seconddigit := c % 16

		z01.PrintRune(hex2Rune(firstdigit))
		z01.PrintRune(hex2Rune(seconddigit))

		z01.PrintRune(' ')

		if i%4 == 3 {
			z01.PrintRune('\n')
		}
	}
	z01.PrintRune('\n')
	for _, c := range arr {
		if c >= 32 && c <= 126 {
			z01.PrintRune(rune(c))
		} else {
			z01.PrintRune(rune('.'))
		}
	}

}
