package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	str := os.Args[1]
	wordsInBetween := false
	for i, c := range str {
		if c == ' ' || c == '\t' {
			if wordsInBetween && i != 0 && str[i-1] != ' ' && str[i-1] != '\t' {
				z01.PrintRune(' ')
			}
		} else {
			z01.PrintRune(c)
			wordsInBetween = true
		}
	}
	z01.PrintRune('\n')
}
