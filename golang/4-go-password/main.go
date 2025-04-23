package main

import (
	"fmt"
	"strings"
	"demo/password/account"
	"demo/password/files"
	"demo/password/output"
	"github.com/fatih/color"
)


var menu = map[string]func(*account.VaultWithDb){
    "1": createAccount,
    "2": searchAccountByUrl,
    "3": searchAccountByLogin,
    "4": deleteAccount,
}


var menuVariants = []string{
    "1. Create account",
    "2. Search account by URL",
    "3. Search account by Login",
    "4. Remove account",
    "5. Exit",
    "Choose variant: ",
}


func menuCounter() func() {
    i := 0
    return func() {
        i++
        fmt.Println(i)
    }
}


func main() {
    // 1. Create account
    // 2. Search account by URL
    // 3. Search account by Login
    // 4 Remove account
    // 5. Exit
    fmt.Println("__Password manager__")
    vault := account.NewVault(files.NewJsonDb("data.json"))
    counter := menuCounter()
Menu:
    for {
        counter()
        variant := promptData(menuVariants...)
        menuFunc := menu[variant]
        if menuFunc == nil {
            break Menu
        }
        menuFunc(vault)
    }
}


func searchAccountByUrl(vault *account.VaultWithDb) {
    // URL
    url := promptData("Input URL for search")
    // Search
    accounts := vault.FindAccounts(url, func(acc account.Account, str string)bool {
        return strings.Contains(acc.Url, str)
    })
    // Output
    outputResult(&accounts)
}


func searchAccountByLogin(vault *account.VaultWithDb) {
    // Login
    login := promptData("Input Login for search")
    // Search
    accounts := vault.FindAccounts(login, func(acc account.Account, str string)bool {
        return strings.Contains(acc.Login, str)
    })
    // Output
    outputResult(&accounts)
}


func outputResult(accounts *[]account.Account) {
    if len(*accounts) == 0 {
        output.PrintError("Accounts Not Found")
    }
    for _, account := range *accounts {
        account.Output()
    }
}


func deleteAccount(vault *account.VaultWithDb) {
    // URL
    url := promptData("Input URL for delete")
    // Remove from vault
    isDeleted := vault.DeleteAccountByUrl(url)
    // Check is removed or not
    if isDeleted {
        color.Green("Removed")
    } else {
        output.PrintError("Not found")
    }
}


func createAccount(vault *account.VaultWithDb) {
    login := promptData("Input login")
	password := promptData("Input password")
	url := promptData("Input url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Invalid url format")
		return
	}
    vault.AddAccount(*myAccount)
}


// Func receive slice of any type
func promptData(prompt ...string) string {
    for i, line := range prompt {
        if i == len(prompt)-1 {
	        fmt.Printf("%v: ", line)
        } else {
            fmt.Println(line)
        }
    }
	var res string
	fmt.Scan(&res)
	return res
}
