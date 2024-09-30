package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	result := len(str) / 2
	return str[:result]
}

func main() {
	table := []string{"hello", "", "Kimetsu no Yaiba", "Z", "123@live.fr", "write ==> 45m$", "-552", "This is the 1st halfThis is the 2nd half", "Hello World"}
	for _, s := range table {
		fmt.Println(RetainFirstHalf(s))
	}
}
