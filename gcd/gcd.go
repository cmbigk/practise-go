package main

import "fmt"

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(Gcd(425, 100))
	fmt.Println(Gcd(429, 127))
	fmt.Println(Gcd(29871, 2500))
	fmt.Println(Gcd(23, 4))
	fmt.Println(Gcd(45, 0))
}
