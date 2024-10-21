package main

import "fmt"

func LastWord(s string) string {
	result := ""
	startChar := false

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !startChar {
			startChar = true
		}

		if startChar {
			if s[i] == ' ' {
				break
			}
			result = string(s[i]) + result
		}
	}
	return result
}

func main() {
	fmt.Println(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(LastWord(" lorem,ipsum "))
	fmt.Println(LastWord(" "))
}
