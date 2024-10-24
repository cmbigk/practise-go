package main

import "fmt"

func main() {
	result := ""
	for i := '9'; i >= '0'; i-- {
		for j := i - 1; j >= '0'; j-- {
			for k := j - 1; k >= '0'; k-- {
				if result != "" {
					result += ", "
				}
				result += fmt.Sprintf("%c%c%c", i, j, k)
			}

		}

	}
	fmt.Print(result)
}
