package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Open the input file for reading
	inputFile, err := os.Open("haha.txt")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a new output file for writing
	outputFile, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	// Read each line from the input file and write it to the output file
	for scanner.Scan() {
		line := scanner.Text()
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to output file:", err)
			return
		}
	}

	// Check for any errors while reading the input file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
	}

	// Flush the buffered data to the output file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing output file:", err)
	}
}
