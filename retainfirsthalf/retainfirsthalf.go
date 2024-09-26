package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	spiltHalf := len(str) / 2
	return str[:spiltHalf]
}
func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half")) // Output: This is the 1st half
	fmt.Println(RetainFirstHalf("A"))                                        // Output: A
	fmt.Println(RetainFirstHalf(""))                                         // Output: ""
	fmt.Println(RetainFirstHalf("Hello World"))
	fmt.Println(RetainFirstHalf("Kimetsu no Yaiba"))
	fmt.Println(RetainFirstHalf("123@live.fr"))
	fmt.Println(RetainFirstHalf("-552"))
}
