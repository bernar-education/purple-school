package main

import "fmt"

// func main() {
// 	a := [4]int{1, 2, 3, 4}
// 	reverse(&a)
// 	fmt.Println(a)
// }

// func reverse(arr *[4]int) {
// 	start, end := 0, len(*arr) - 1
// 	for start < end {
// 		(*arr)[start], (*arr)[end] = (*arr)[end], (*arr)[start]
// 		start++
// 		end--
// 	}
// }

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Input login")
	password := promptData("Input password")
	url := promptData("Input url")

	myAccount := account{
		password: password,
		url:      url,
		login:   login,
	}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc)
	fmt.Println(acc.login, acc.password, acc.url)
}
