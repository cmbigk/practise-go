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

func removePunctuation(pun []string) []string {

	// Ensure pun is not empty before checking pun[0]
	if len(pun) == 0 {
		return pun
	}

	// Handle punctuation at the start of the first word (index 0)
	for len(pun[0]) > 0 && strings.Contains(".,!?;:", string(pun[0][0])) {
		// If the first character of the first word is punctuation, remove it
		pun[0] = pun[0][1:]
	}

	// Process all other words starting from index 1
	for i := 1; i < len(pun); i++ {
		// If the current word starts with punctuation
		for len(pun[i]) > 0 && strings.Contains(".,!?;:", string(pun[i][0])) {
			// Attach punctuation to the previous word
			pun[i-1] = pun[i-1] + string(pun[i][0])
			// Remove the punctuation from the current word
			pun[i] = pun[i][1:]
		}
	}

	return removeTheSlice(pun)
}

func main() {
	input := []string{"Hello", ",", "world", "!", "How", "are", "you", "?"}
	fmt.Println("Original:", input)

	output := removePunctuation(input)
	fmt.Println("Processed:", output)

	testCases := [][]string{
		{"Hello", ",", "world", "!"},
		{"!!!", "How", "are", "you?", "?"},
		{".....", ".", ".", "Test"},
		{"NoPunctuationsHere"},
	}

	for _, testCase := range testCases {
		fmt.Println("Testing:", testCase)
		result := removePunctuation(testCase)
		fmt.Println("Result:", result)
	}
}
