package main

import (
	"fmt"
)

func main() {
	fmt.Println(IsCapitalized("Hello! How are you?"))
	fmt.Println(IsCapitalized("Hello How Are You"))
	fmt.Println(IsCapitalized("Whats 4this 100K?"))
	fmt.Println(IsCapitalized("Whatsthis4"))
	fmt.Println(IsCapitalized("!!!!Whatsthis4"))
	fmt.Println(IsCapitalized(""))
}

func isLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i <= len(s); i++ {
		if i > 0 && s[i-1] == ' ' && isLower(s[i]) {
			return false
		}
		if isLower(s[0]) {
			return false
		}

	}
	return true

}
