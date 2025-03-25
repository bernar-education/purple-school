package main

import "fmt"

// В цикле спрашиваем ввод транзакции: 10, 20, 30, 40
// Добавлять каждую транзакцию в массим транзакции
// Вывести массив
// !Вывести сумму баланса в консоль

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
	fmt.Printf("Сумма всех транзакций: %.2f", balance)
}

func getInputTransaction() float64 {
	var transaction float64
	fmt.Print("Введите значение (для выхода введите n): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(transactions []float64) (balance float64) {
	for _, value := range transactions {
		balance += value
	}
	return
}
