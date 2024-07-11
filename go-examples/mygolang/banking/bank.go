package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	os.WriteFile(accountBalanceFileName, []byte(fmt.Sprintf("%f", balance)), 0644)

}

func readBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFileName)
	if err != nil {
		return 1000, errors.New("failed to read balance file")
	}

	balance, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("failed to parse balance file")
	}
	return balance, nil
}
func main() {

	var accountBalance, err = readBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("--------")
		panic(err)
	}

	fmt.Println("Welcome to the Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check your balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Enter the amount you want to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount, must be greater than 0")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Println("Enter the amount you want to withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("Insufficient balance")
				continue
			}
			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is: ", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Thank you for using the Bank!")
			return
		default:
			fmt.Println("Invalid choice")
		}

	}

}
