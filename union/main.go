package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isContain(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}

func main() {
	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	result := ""
	str1 := os.Args[1]
	str2 := os.Args[2]

	for _, char := range str1 {
		if !isContain(result, char) {
			result += string(char)
		}
	}

	for _, char := range str2 {
		if !isContain(result, char) {
			result += string(char)
		}
	}

	for _, char := range result {
		z01.PrintRune(char)
	}

	z01.PrintRune('\n')
}
