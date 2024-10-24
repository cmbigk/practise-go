package main

import (
	"fmt"
	"strings"
)

func removeTheSlice(remove []string) []string {

	var result []string

	for _, char := range remove {
		if char != "" {
			result = append(result, char)
		}

	}
	return result
}

func addAn(vowel []string) []string {
	for i := 0; i < len(vowel)-1; i++ {
		// Check if the current word is "a"
		if vowel[i] == "a" && len(vowel[i+1]) > 0 && strings.Contains("aeiouh", string(vowel[i+1][0])) {
			// Replace "a" with "an" if the next word starts with a vowel or 'h'
			vowel[i] = "an"
		}
	}
	return removeTheSlice(vowel)
}

func main() {
	// Test cases
	test1 := []string{"This", "is", "a", "apple"}
	test2 := []string{"I", "have", "a", "house", "and", "a", "umbrella"}
	test3 := []string{"He", "is", "a", "honest", "man"}
	test4 := []string{"This", "is", "not", "an", "empty", "", "sentence"}

	// Testing addAn function
	result1 := addAn(test1)
	result2 := addAn(test2)
	result3 := addAn(test3)
	result4 := addAn(test4)

	// Print results
	fmt.Println("Test 1:", result1)
	fmt.Println("Test 2:", result2)
	fmt.Println("Test 3:", result3)
	fmt.Println("Test 4:", result4)
}
