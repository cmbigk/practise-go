package main

func DigitLen(n, base int) int {
	if base < 2 && base > 36 {
		return -1
	}
	if n < 0 {
		n = -n
	}
	counter := 0
	for n > 0 {
		n = n / base
	}
	return counter
}
