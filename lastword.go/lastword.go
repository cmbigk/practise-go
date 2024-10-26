package main

import (
	"fmt"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}

func LastWord(s string) string {
	var lastword string
	foundWord := false

	for _, char := range s {
		if char == ' ' {
			if !foundWord {
				continue
			}
			if foundWord {
				break
			}

		}
		lastword = lastword + string(char)
		foundWord = true

	}
	return lastword + "\n"
}
