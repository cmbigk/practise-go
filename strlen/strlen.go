package main

import "fmt"

func StrLen(s string) int {
	return len([]rune(s))
}
func main() {
	l := StrLen("richarlison!")
	fmt.Println(l)
}
