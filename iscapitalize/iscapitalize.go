package main

import "fmt"

// IsCapitalized checks if all words in the string are capitalized.
func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	for i := 0; i < len(s); i++ {
		if i > 0 && s[i-1] == ' ' && s[i] >= 'a' && s[i] <= 'z' {
			return false
		}
	}

	if s[0] >= 'a' && s[0] <= 'z' {
		return false
	}
	return true
}

func main() {
	// Define some test cases
	testCases := []struct {
		input    string
		expected bool
	}{
		{"Hello! €How are you?", false},
		{"A Simple Sentence", true},
		{"a lowercase start", false},
		{"", false},
		{"HELLO WORLD", true},
		{"H3110 W0r1d!", true},
		{"Hello How Are You", true},
		{"What's This?", true},
		{"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890", false},
	}

	// Iterate over each test case and check the output of IsCapitalized
	for _, testCase := range testCases {
		result := IsCapitalized(testCase.input)
		fmt.Printf("Input: '%s' | Expected: %v | Result: %v\n", testCase.input, testCase.expected, result)
	}
}
