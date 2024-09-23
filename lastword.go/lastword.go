package piscine

func LastWord(s string) string {
	var word string
	started := false

	// Loop from the end of the string to the beginning
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' && !started {
			started = true // We start reading the last word
		}
		if started {
			if s[i] == ' ' {
				break // Stop when we hit the first space before the last word
			}
			word = string(s[i]) + word // Prepend characters to the word
		}
	}
	return word + "\n" // Add newline after the last word
}
