package main

import "fmt"

func main() {
	var accountBalance float64 = 16648.85
	fmt.Println("Welcome to BIGK Bank!")

	for {

		fmt.Println("What service do we need to offer, Sir?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("The deposit amount cannot be less than 0$")
				continue
			}
			newaccountBalance := accountBalance + depositAmount
			fmt.Println("Your new balance is: ", newaccountBalance)

		case 3:
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount <= 0 {
				fmt.Println("The withdraw amount cannot be less than 0$")
				continue
			}
			if withdrawAmount > accountBalance {
				fmt.Println("Not enough balance")
				continue
			}

			newaccountBalance := accountBalance - withdrawAmount
			fmt.Println("Your new balance is: ", newaccountBalance)

		default:
			fmt.Println("Thank you for using BIGK Bank!, We make sure to provide the best service to you!")

		}
	}
	fmt.Println("TQ for using our service!")
}
