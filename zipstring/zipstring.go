package main

import (
	"fmt"
	"strconv"
)

func ZipString(s string) string {
	result := ""
	i := 0

	for i < len(s) {
		count := 1
		for i+1 < len(s) && s[i] == s[i+1] {
			count++
			i++
		}
		result += strconv.Itoa(count) + string(s[i])
		i++
	}
	return result
}

func main() {
	args := []string{
		"aaaa",
		"zzzzzZZZZZZ",
		"",
		"ziiiiipMeee",
		"hhellloTthereYouuunggFelllas",
	}

	for _, arg := range args {
		fmt.Println(ZipString(arg))
	}
}
