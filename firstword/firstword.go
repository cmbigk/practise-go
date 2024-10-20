package main

import "fmt"

/*
Write a function that takes a string as input and returns the first word of the string.
If the input string is empty or contains only whitespace characters, the function should return an empty string.
*/

func FirstWord(s string) string {
	result := ""
	wordStarted := false
	for _, char := range s {
		if char != ' ' && !wordStarted {
			wordStarted = true
		}
		if wordStarted {
			if char == ' ' {
				break

			}
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(FirstWord("Hello World!"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("   salut !!! "))
	fmt.Println(FirstWord("Hello World!"))
	fmt.Println(FirstWord("  afee dfw!"))

}
