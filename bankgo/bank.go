package main

import (
	"fmt"
	"os"

	"bankexample.com/displayoptions"
	"github.com/Pallinder/go-randomdata"
)

func writeLogsOnfile(name string, balance float64) {
	balancetext := fmt.Sprintf(" %s has %.2f balance ", name, balance)
	err := os.WriteFile("balance.txt", []byte(balancetext), 0o644)
	if err != nil {
		fmt.Println("An Error occured ...........")
	}
}

func main() {
	var name string = "meow"
	var accountBalance float64 = 10000.00
	var choice int32
	fmt.Println(randomdata.Address())
	writeLogsOnfile(name, accountBalance)
	for choice != 4 {
		displayoptions.DisplayOptions()
		fmt.Println("Please Select Your choice :")
		fmt.Scan(&choice)
		fmt.Printf("You Have Selectd %d \n", choice)
		switch choice {

		case 1:
			fmt.Printf("Your Balance is : %.2f \n", accountBalance)
		case 2:
			fmt.Println("Enter Money You wanna deposit")
			var moneyToBeDeposited float64
			fmt.Scan(&moneyToBeDeposited)
			if moneyToBeDeposited > 0 {
				accountBalance += moneyToBeDeposited
				writeLogsOnfile(name, accountBalance)
				fmt.Printf("New Balanace is : %.3f", accountBalance)
			} else {
				fmt.Println("Money to be Deposited Should be Greater Than Zero !")
			}

		case 3:
			var withdrawAmount float64
			fmt.Println("Enter Money You Wnna Withdraw !")
			fmt.Scan(&withdrawAmount)
			if withdrawAmount > accountBalance {
				fmt.Println("You Dont Have Enough Amount !")
			} else {
				accountBalance -= withdrawAmount
				writeLogsOnfile(name, accountBalance)
				fmt.Printf("Succesfull and New AMount Balance is : %.2f", accountBalance)
			}
		default:
			fmt.Println("Thank you !")
			return
		}
		fmt.Println("******** Do You Wanna Continue! (y/n) **********")
		var userInput string
		fmt.Scan(&userInput)

		if userInput == "y" {
			continue
		} else if userInput == "n" {
			fmt.Println("Exiting...")
			break
		} else {
			fmt.Println("Invalid input, please enter y or n")
		}
	}
}
