package main

import (
	"fmt"
)

func main() {
	initialPrice := 100.0
	sellingPrice := 120.0
	discountRate := 2.0

	fmt.Print("Initial Price: ")
	fmt.Scan(&initialPrice)

	fmt.Print("Selling Price: ")
	fmt.Scan(&sellingPrice)

	fmt.Print("Discount Rate: ")
	fmt.Scan(&discountRate)

	realPrice := sellingPrice * (1 - discountRate/100)
	profit := realPrice - initialPrice
	fmt.Println("Real Price: ", realPrice)
	fmt.Println("Profit: ", profit)
}
