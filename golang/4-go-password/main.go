package main

import (
	"demo/password/account"
	"demo/password/files"
	"fmt"
)


func main() {
    // 1. Create account
    // 2. Search account
    // 3. Remove account
    // 4. Exit
    fmt.Println("__Password manager__")
Menu:
    for {
        variant := getMenu()
        switch variant {
        case 1:
            createAccount()
        case 2:
            searchAccount()
        case 3:
            deleteAccount()
        default:
            break Menu
        }
    }
}

func getMenu() int {
    var variant int
    fmt.Println("Choose variant: ")
    fmt.Println("1. Create account")
    fmt.Println("2. Search account")
    fmt.Println("3. Remove account")
    fmt.Println("4. Exit")
    fmt.Scan(&variant)
    return variant
}


func searchAccount() {

}


func deleteAccount() {

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
    vault := account.NewVault()
    vault.AddAccount(*myAccount)
    data, err := vault.ToByteSlice()
	if err != nil {
		fmt.Println("Can't convert to byte slice")
		return
	}
    files.WriteToFile(data, "data.json")
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
