package main

import (
	"fmt"
)

func UltimateDivMod(a *int, b *int) {
	p := *a
	q := *b
	*a = p / q
	*b = p % q
}
func main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
