package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	firstStr := os.Args[1]
	secondStr := os.Args[2]

	if len(os.Args) != 3 {
		z01.PrintRune(' ')
		return
	}

	i := 0
	for _, p := range secondStr {
		if p == rune(firstStr[i]) {
			i++
		}

		if i == len(firstStr) {
			break
		}
	}
	if i == len(firstStr) {
		for _, v := range firstStr {
			z01.PrintRune(v)
		}
		z01.PrintRune('\n')
	}
}
