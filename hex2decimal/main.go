package main

import "fmt"

func hex2decimal(s rune) int {
	if s >= 'A' && s <= 'F' {
		return int(s - 'A' + 10)
	}

	if s >= 'a' && s <= 'f' {
		return int(s - 'a' + 10)
	}

	if s >= '0' && s <= '9' {
		return int(s - '0')
	}

	return -1
}
func main() {
	fmt.Println(hex2decimal('9'))
	fmt.Println(hex2decimal('A'))
	fmt.Println(hex2decimal('F'))
	fmt.Println(hex2decimal('a'))
	fmt.Println(hex2decimal('f'))
	fmt.Println(hex2decimal('G'))
}
