package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)


func main() {
    createAccount()
}

func createAccount() {
    login := promptData("Input login")
	password := promptData("Input password")
	url := promptData("Input url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Invalid url format")
		return
	}
	file, err :=myAccount.ToByteSlice()
	if err != nil {
		fmt.Println("Can't convert to byte slice")
		return
	}
    files.WriteToFile(file, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
