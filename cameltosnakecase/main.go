package main

import "fmt"

func isUpper(s rune) bool {
	return s >= 'A' && s <= 'Z'
}

func isLower(s rune) bool {
	return s >= 'a' && s <= 'z'
}
func CamelToSnakeCase(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if isLower(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			result += string(s[i])
		} else if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && isLower(rune(s[i+1])) {
			result += "_"
			result += string(s[i])
		} else {
			return s
		}

	}
	return result
}

func main() {
	args := []string{
		"CamelCase",
		"CAMELCase",
		"HelloWorld",
		"132",
		" ",
		"",
		"A",
		"abcs",
		"AbC",
		"AbCEf",
		"abcAree",
		"ahe1Abde",
		"TesTing",
		"SOMEVARIABLE",
		"ASuperLonGVariableName",
		"thisIsaTestOfCamelCase",
		"aA",
	}
	for _, arg := range args {
		fmt.Println(CamelToSnakeCase(arg))
	}
}
