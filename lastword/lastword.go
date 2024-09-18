package main

import (
	"fmt"
	"strings"
)

func LastWord(s string) string {
	checkSpace := strings.TrimSpace(s)
	if checkSpace == "" {
		return "\n"
	}
	result := strings.Fields(checkSpace)
	if len(result) == 0 {
		return "\n"
	}
	return result[len(result)-1] + "\n"
}
func main() {
	fmt.Println(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Println(LastWord(" lorem,ipsum "))
	fmt.Println(LastWord(" "))
}
