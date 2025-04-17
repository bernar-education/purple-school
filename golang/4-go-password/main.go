package main

import (
	"demo/password/account"
	"fmt"
	"demo/password/files"
	"demo/password/output"
	"github.com/fatih/color"
)


func main() {
    // 1. Create account
    // 2. Search account
    // 3. Remove account
    // 4. Exit
    fmt.Println("__Password manager__")
    // vault := account.NewVault(cloud.NewCloudDb("https://berber.com"))
    vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
    for {
        variant := promptData([]string{
            "1. Create account",
            "2. Search account",
            "3. Remove account",
            "4. Exit",
            "Choose variant: ",
        })
        switch variant {
        case "1":
            createAccount(vault)
        case "2":
            searchAccount(vault)
        case "3":
            deleteAccount(vault)
        default:
            break Menu
        }
    }
}


func searchAccount(vault *account.VaultWithDb) {
    // URL
    url := promptData([]string{"Input URL for search"})
    // Search
    accounts := vault.FindAccountByUrl(url)
    if len(accounts) == 0 {
        output.PrintError("Accounts Not Found")
    }
    // Output
    for _, account := range accounts {
        account.Output()
    }
}


func deleteAccount(vault *account.VaultWithDb) {
    // URL
    url := promptData([]string{"Input URL for delete"})
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
    login := promptData([]string{"Input login"})
	password := promptData([]string{"Input password"})
	url := promptData([]string{"Input url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		output.PrintError("Invalid url format")
		return
	}
    vault.AddAccount(*myAccount)
}


// Func receive slice of any type
func promptData[T any](prompt []T) string {
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
