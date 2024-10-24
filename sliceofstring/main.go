package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"}
	s[0] = "d"

	fmt.Println(s)
}
