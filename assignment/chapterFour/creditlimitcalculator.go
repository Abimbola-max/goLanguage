package main

import "fmt"

func main() {
	fmt.Println("Enter '-1' as the account number to stop.")

	for {
		accountNumber := accountNumberInput()
		if accountNumber == -1 {
			break
		}

		balance := balanceAtTheBeginning()
		totalCharged := totalItemCharged()
		totalCredits := totalCreditsInput()
		allowedCreditLimit := allowedCredit()

		newBalance := calculateNewBalance(balance, totalCharged, totalCredits)

		fmt.Printf("Account Number: %d\n", accountNumber)
		fmt.Printf("New balance: %d\n", newBalance)

		if newBalance > allowedCreditLimit {
			fmt.Println("Credit limit exceeded")
		} else {
			fmt.Println("Credit limit not exceeded")
		}
		fmt.Println()
	}

	fmt.Println("Thank you for using the Credit Limit Calculator.")

}

func calculateNewBalance(balanceAtTheBeginning int, totalItemCharged int, totalCredits int) int {
	newBalance := balanceAtTheBeginning + totalItemCharged - totalCredits
	return newBalance
}

func accountNumberInput() int {
	var accountNumber int
	fmt.Println("Enter account number: ")
	fmt.Scanln(&accountNumber)
	return accountNumber
}

func balanceAtTheBeginning() int {
	var balanceAtTheBeginning int
	fmt.Println("Enter balance at the beginning of the month: ")
	fmt.Scanln(&balanceAtTheBeginning)
	return balanceAtTheBeginning
}

func totalItemCharged() int {
	var totalItemCharged int
	fmt.Println("Enter total of all items charged this month")
	fmt.Scanln(&totalItemCharged)
	return totalItemCharged
}

func totalCreditsInput() int {
	var totalCredits int
	fmt.Println("Enter total credits: ")
	fmt.Scanln(&totalCredits)
	return totalCredits
}

func allowedCredit() int {
	var allowedCreditLimit int
	fmt.Print("Enter allowed credit limit: ")
	fmt.Scanln(&allowedCreditLimit)
	return allowedCreditLimit
}
