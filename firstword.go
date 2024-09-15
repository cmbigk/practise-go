package main

import "strings"

func FirstWord(s string) string {
	checkSpaces := strings.TrimSpace(s)
	if checkSpaces == "" {
		return ""
	}
	result := strings.Split(s, " ")
	return result[0]
}
