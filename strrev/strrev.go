package main

import "fmt"

func StrRev(s string) string {
	akm := ""
	for v := len(s) - 1; v >= 0; v-- {
		akm += (string(s[v]))
	}
	return akm
}
func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}
