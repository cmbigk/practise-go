package main

import "fmt"

func PrintRevComb() {
	result := ""
	for c1 := '9'; c1 >= '0'; c1-- {
		for c2 := c1 - 1; c2 >= '0'; c2-- {
			for c3 := c2 - 1; c3 >= '0'; c3-- {
				if result != "" {
					result += ", "
				}
				result += fmt.Sprintf("%c%c%c", c1, c2, c3)
			}
		}
	}
}

func main() {
	fmt.Println(PrintRevComb)
}
