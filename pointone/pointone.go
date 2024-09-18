package main

import "fmt"

func PointOne(n *int) {
	*n = 69
}

func main() {
	a := 0
	PointOne(&a)
	fmt.Println(a)
}
