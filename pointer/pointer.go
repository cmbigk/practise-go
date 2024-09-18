package main

import "fmt"

func main() {
	var a *int
	b := 6969
	a = &b
	fmt.Println(*a)
}
