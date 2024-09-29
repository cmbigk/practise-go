package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	// If the number of arguments is not exactly 1 (excluding the program name)
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	// Input string from the command line argument
	input := os.Args[1]

	// Variables to track space trimming and word processing
	insideWord := false
	for i, c := range input {
		// Skip leading spaces and tabs
		if c == ' ' || c == '\t' {
			if insideWord && i > 0 && (input[i-1] != ' ' && input[i-1] != '\t') {
				z01.PrintRune(' ')
			}
		} else {
			// Print non-space character
			z01.PrintRune(c)
			insideWord = true
		}
	}

	// Print a newline at the end
	z01.PrintRune('\n')
}
