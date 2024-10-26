package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isInStr(str1, str2 string) bool {
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] == str2[j] {
			i++
		}
		j++

	}
	return i == len(str1)
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	str1 := os.Args[1]
	str2 := os.Args[2]

	if isInStr(str1, str2) {
		for _, char := range str1 {
			z01.PrintRune(char)
		}

	}
	z01.PrintRune('\n')
}
