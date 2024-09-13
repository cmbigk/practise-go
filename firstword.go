package main

import (
	"fmt"
	"piscine"
	"strings"
)

func FirstWord(s string) string {
	checkSpace := strings.TrimSpace(s)

	result := strings.Split(checkSpace, " ")
	return result[0]
}

func main() {
	fmt.Print(piscine.FirstWord("hello there"))
	fmt.Print(piscine.FirstWord(""))
	fmt.Print(piscine.FirstWord("hello   .........  bye"))
}
