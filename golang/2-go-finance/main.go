package main

import "fmt"

// In a loop, ask for transaction input: 10, 20, 30, 40
// Add each transaction to the transactions array
// Print the array
// !Print the total balance to the console

func main() {
	transactions := []float64{}
	for {
		transaction := getInputTransaction()
		if transaction == 0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	balance := calculateBalance(transactions)
	fmt.Printf("Sum of all transactions: %.2f", balance)
}

func getInputTransaction() float64 {
	var transaction float64
	fmt.Print("Input transaction value (for exit enter n): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) (balance float64) {
	for _, value := range transactions {
		balance += value
	}
	return
}
