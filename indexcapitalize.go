package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the input file for reading
	inputFile, err := os.Open("haha.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(inputFile)

	// The word you want to capitalize
	wordToCapitalize := "F"
	targetIndex := 0 // Capitalize the second occurrence (index 1)

	// Modify the specific occurrence of the word
	var modifiedContent []string
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		var modifiedWords []string
		lineOccurrenceCount := 0
		for _, word := range words {
			// Convert the word to lowercase for comparison
			lowerWord := strings.ToLower(word)
			if lowerWord == wordToCapitalize {
				lineOccurrenceCount++
				if lineOccurrenceCount == targetIndex+1 { // Check if it's the target occurrence
					modifiedWords = append(modifiedWords, strings.Title(word)) // Capitalize the word
				} else {
					modifiedWords = append(modifiedWords, strings.ToLower(word)) // Lowercase the word
				}
			} else {
				modifiedWords = append(modifiedWords, strings.ToLower(word)) // Lowercase non-target words
			}
		}
		modifiedContent = append(modifiedContent, strings.Join(modifiedWords, " "))
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
		return
	}

	// Write the modified content to the output file
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	for _, line := range modifiedContent {
		fmt.Fprintln(outputFile, line)
	}

	fmt.Println("File processing completed successfully.")
}
