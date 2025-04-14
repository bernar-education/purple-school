package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)


func main() {
    files.ReadFromFile()
    files.WriteToFile("Привет", "file.txt")
	login := promptData("Input login")
	password := promptData("Input password")
	url := promptData("Input url")

	myAccount, err := account.NewAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid url format")
		return
	}
	myAccount.OutputPassword()

	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
