package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	c := len(str) / 2
	return str[:c]
}
func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
