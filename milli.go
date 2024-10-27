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

	inWord := false
	for i, char := range str {
		if char == ' ' || char == '\t' {
			if inWord && i != 0 && str[i-1] != ' ' && str[i-1] != '\t' {
				z01.PrintRune(' ')
				z01.PrintRune(' ')
				z01.PrintRune(' ')
			}
		} else {
			z01.PrintRune(char)
			inWord = true
		}

	}
	z01.PrintRune('\n')

}
