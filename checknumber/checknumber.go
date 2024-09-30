package main

import "fmt"

func checknumber(k string) bool {
	for _, c := range k {
		if c >= '0' && c <= '9' {
			return true
		}
	}
	return false
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
		fmt.Printf("checknumber(%q) = %t\n", arg, checknumber(arg))
	}
}
