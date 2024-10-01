package main

import "fmt"

/*
Write a function that takes a string as input and returns the first word of the string.
If the input string is empty or contains only whitespace characters, the function should return an empty string.
*/

func FirstWord(s string) string {
	var word string
	startc := false
	for _, c := range s {
		if c != ' ' && !startc {
			startc = true
		}
		if startc {
			if c == ' ' {
				break
			}
			word += string(c)
		}
	}
	return word
}
func main() {
	fmt.Println(FirstWord("Hello World!"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("   salut !!! "))
	fmt.Println(FirstWord("Hello World!"))
	fmt.Println(FirstWord("  afee dfw!"))

}
