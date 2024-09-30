package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return // Do nothing if the number of arguments is not 1
	}

	str := os.Args[1]
	firstWord := true // Flag to check if we're printing the first word

	inWord := false
	for i, r := range str {
		// Check for spaces or tabs
		if r == ' ' || r == '\t' {
			if inWord { // If we were inside a word
				inWord = false // We're now out of a word
				if i > 0 && (str[i-1] != ' ' && str[i-1] != '\t') {
					// Print three spaces if not the first word
					if !firstWord {
						z01.PrintRune(' ')
						z01.PrintRune(' ')
						z01.PrintRune(' ')
					}
				}
			}
		} else {
			// Print the character
			z01.PrintRune(r)
			inWord = true
			firstWord = false // We have printed the first word
		}
	}

	z01.PrintRune('\n') // Print a newline at the end
}
