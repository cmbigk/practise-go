package main

import "fmt"

func printCalculation(a int, b int) (int, int) {
	var result, reminder int
	result = a / b
	reminder = a % b
	return result, reminder
}
func main() {
	a := 3658291658713951356
	b := 4543553
	result, reminder := printCalculation(a, b)
	fmt.Printf("The result is %v and the reminder is %v", result, reminder)
}
