package main

import (
	"fmt"
)

func main() {
	initialPrice := 100.0
	discountRate := 2.0

	fmt.Print("Initial Price: ")
	fmt.Scan(&initialPrice)

	fmt.Print("Discount Rate: ")
	fmt.Scan(&discountRate)

	realPrice := initialPrice * (1 - discountRate/100)
	fmt.Println("Real Price: ", realPrice)

}
