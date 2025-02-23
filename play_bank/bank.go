package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"onis-emem.com/go/play-bank/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("Oops! Something happened")
		fmt.Println(err)
		fmt.Println("----------------------")
		panic("Can't proceed at the moment")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println(randomdata.PhoneNumber())

	for {
		presentBankingOptions()

		var choice int
		fmt.Print("Choose for from number 1 to 4: ")
		fmt.Scan(&choice)

		// you can assign a condition to a variable
		// wantsCheckBalance := choice == 1

		switch choice {
		case 1:
			fmt.Println("Your account balance is: ", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			// Check if deposit is greater than 0
			if depositAmount <= 0 {
				fmt.Println("Invalid amount. Deposit must be greater than 0.")

				// Return will exit so we change to continue key word
				// return
				continue
			}

			accountBalance += depositAmount // accountBalance = accountBalance + depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your Withdrawal amount: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. Amount must be greater than 0.")
				//return
				continue
			}

			if withdrawAmount > accountBalance {
				fmt.Println("You have insufficient balance")
				return
			}

			accountBalance -= withdrawAmount // accountBalance = accountBalance - withdrawAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Print("Goodbye!")
			fmt.Println("Thank you for choosing Go Bank")
			return
			// break would not work in switch case.
		}

	}
}
