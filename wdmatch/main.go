package main

import (
	"os"

	"github.com/01-edu/z01"
)

func IsDatHere(c string, d rune) bool {
	for _, v := range c {
		if v == d {
			return true
		}
	}
	return false
}

func main() {
	firstStr := os.Args[1]
	secondStr := os.Args[2]

	if len(os.Args) != 3 {
		z01.PrintRune(' ')
		return
	}

	i := 0
	for _, char := range secondStr {
		if char == rune(firstStr[i]) {
			i++
		}

		if i == len(firstStr) {
			break
		}
	}
	if i == len(firstStr) {
		for _, r := range firstStr {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}
