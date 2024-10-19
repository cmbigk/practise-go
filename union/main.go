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
		z01.PrintRune('\n')
	}

	var answer string
	for _, v := range firstStr {
		if !IsDatHere(answer, v) {
			answer += string(v)
		}
	}

	for _, v := range secondStr {
		if !IsDatHere(answer, v) {
			answer += string(v)
		}
	}

	for _, r := range answer {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}
