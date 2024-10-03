package main

import "fmt"

var accountBalance float64 = 16648.85

func main() {
	for i := 0; i <= 600; i++ {
		fmt.Println("Welcome to BIGK Bank!")
		fmt.Println("What service do we need to offer, Sir?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Println("Your choice: ")
		fmt.Scan(&choice)
		fmt.Println("You chose: ", choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", accountBalance)
			fmt.Println("Thank you for using BIGK Bank!")
		} else if choice == 2 {
			fmt.Println("How much do you want to deposit?")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("The deposit amount cannot be less than 0$")
				return
			}
			newaccountBalance := accountBalance + depositAmount
			fmt.Println("Your new balance is: ", newaccountBalance)
			fmt.Println("Thank you for using BIGK Bank!")
		} else if choice == 3 {
			fmt.Println("How much do you want to withdraw?")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("The withdraw amount cannot be less than 0$")
				return
			}

			if withdrawAmount > accountBalance {
				fmt.Println("Not enough balance")
				return
			}

			newaccountBalance := accountBalance - withdrawAmount
			fmt.Println("Your new balance is: ", newaccountBalance)
			fmt.Println("Thank you for using BIGK Bank!")
		} else if choice == 4 {
			fmt.Println("Thank you for using BIGK Bank!, We make sure to provide the best service to you!")
		} else {
			fmt.Println("Choose a valid option, Thank you!")
		}
		break
	}
	fmt.Println("TQ for using our service!")
}
