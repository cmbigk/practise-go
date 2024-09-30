package main

import "fmt"

func CountAlpha(s string) int {
	count := 0
	for _, k := range s {
		if (k >= 'a' && k <= 'z') || (k >= 'A' && k <= 'Z') {
			count++
		}
	}
	return count
}

func main() {
	args := []string{
		"123",
		"H1ll0",
		"",
		"1",
		"1.1",
		"Containe1number",
		"     ",
		"upson lorem ipsum",
	}

	for _, arg := range args {
		fmt.Printf("CountAlpha(%q) = %d\n", arg, CountAlpha(arg))
	}
}
