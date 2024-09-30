package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		return
	}
	str := os.Args[1]
	old := os.Args[2]
	new := os.Args[3]

	// Ensure 'old' and 'new' are single characters
	if len(old) != 1 || len(new) != 1 {
		return
	}

	// Replace occurrences of 'old' with 'new'
	result := ""
	for _, char := range str {
		if string(char) == old { // Compare current character with 'old'
			result += new // Append 'new' if it matches
		} else {
			result += string(char) // Append original character if it doesn't match
		}
	}

	fmt.Println(result)
}
