package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isLower(c rune) bool {
	return c >= 'a' && c <= 'z'
}

func isUpper(c rune) bool {
	return c >= 'A' && c <= 'Z'
}

func reverseStrCap(s string) string {
	result := []rune(s)

	for i := 0; i < len(result); i++ {
		if i == len(result)-1 || result[i+1] == ' ' {
			if isLower(result[i]) {
				result[i] -= 32
			}
		} else if isUpper(result[i]) {
			result[i] += 32
		}

	}
	return string(result)
}

func main() {
	if len(os.Args) > 1 {
		for _, char := range os.Args[1:] {
			result := reverseStrCap(char)

			for _, char := range result {
				z01.PrintRune(char)
			}
			z01.PrintRune('\n')
		}

	}

}
