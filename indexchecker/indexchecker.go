package main

import (
	"fmt"
	"strings"
)

func main() {
	// The provided text to check
	text := `IN THE CONTEXT FILE, ALL LOOK-AND-FEEL SPECIFICATIONS ARE DEFINED BY THE LOOKANDFEEL ELEMENT UNDER CUSTOMSEARCHENGINE. THIS ELEMENT DETERMINES WHETHER ADS ARE DISPLAYED, HOW THE SEARCH RESULTS SECTION ARE RENDERED, AND HOW INDIVIDUAL SEARCH RESULTS ARE DISPLAYED. THE FOLLOWING EXAMPLE SHOWS ALL THE ATTRIBUTES AND CHILD ELEMENTS OF THE LOOKANDFEEL ELEMENT.
THIS PAGE DESCRIBES HOW TO CUSTOMIZE THE APPEARANCE OF YOUR SEARCH ENGINE USING THE CONTEXT FILE, WHICH IS THE XML SPECIFICATION FOR YOUR SEARCH ENGINE.
THE RESULTS ELEMENT CONTROLS THE COLOR OF INDIVIDUAL RESULTS IN THE SEARCH ELEMENT. EACH INDIVIDUAL RESULT FORMS A UNIT OF TITLE, RESULT SNIPPET, AND LINK. DEFINING THIS CHILD ELEMENT LETS YOU VISUALLY DELINEATE INDIVIDUAL RESULTS OR HIGHLIGHT RESULTS THAT ARE BEING SELECTED BY USERS.
THIS ELEMENT CONTROLS THE COLOR FOR THE INDIVIDUAL RESULTS.`

	// Keywords to check for indexing
	keywords := []string{
		"F",
	}

	// Check for each keyword in the text
	for _, keyword := range keywords {
		// Convert the text to lowercase for case-insensitive search
		lowerText := strings.ToLower(text)
		lowerKeyword := strings.ToLower(keyword)

		// Find the index positions of the keyword
		startIndex := 0
		for {
			index := strings.Index(lowerText[startIndex:], lowerKeyword)
			if index == -1 {
				break // No more occurrences found
			}
			// Calculate the actual index in the original text
			actualIndex := startIndex + index
			fmt.Printf("The keyword '%s' is found at index %d.\n", keyword, actualIndex)
			startIndex = actualIndex + 1 // Move past the last found index
		}
	}
}
