package main

import "fmt"

func checkAlphabet(s string) bool {
	for _, c := range s {
		if c < 'a' || c > 'z' {
			return false
		}
	}
	return true
}
func isUpper(s rune) bool {
	if s >= 'A' && s <= 'Z' {
		return true
	}
	return false
}
func isLower(s rune) bool {
	if s >= 'a' && s <= 'z' {
		return true
	}
	return false
}

func CamelToSnakeCase(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if i != 0 && isUpper(rune(s[i])) && i+1 < len(s) && isLower(rune(s[i+1])) {
			result += "_"
			result += string(rune(s[i]))
		} else if isLower(rune(s[i])) || (i == 0 && isUpper(rune(s[i]))) {
			result += string(rune(s[i]))
		} else {
			return s
		}
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
