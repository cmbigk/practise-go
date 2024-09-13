package main

import (
	"fmt"
	"strings"
)

func FirstWord(s string) string {
	checkSpace := strings.TrimSpace(s)
	if checkSpace == "" {
		return "" + "\n"

	}

	result := strings.Split(checkSpace, " ")
	return result[0]
}

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}
