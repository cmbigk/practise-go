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
	str1 := os.Args[1]
	str2 := os.Args[2]

	if len(os.Args) != 3 {
		z01.PrintRune('\n')
		return
	}

	var answer string
	for _, char := range str1 {
		if !isContain(answer, char) {
			answer += string(char)
		}
	}

	for _, char := range str2 {
		if !isContain(answer, char) {
			answer += string(char)
		}
	}

	for _, r := range answer {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
