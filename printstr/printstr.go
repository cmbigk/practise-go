package main

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, w := range s {
		z01.PrintRune(w)

	}
}
func main() {
	PrintStr("gae Richard!")
}
