package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
	"time"
)

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

type accountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	account
}

func (acc *account) outputPassword() {
	fmt.Println(acc.login, acc.password, acc.url)
}

func (acc *account) generatePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

// 1. If login is empty, return error
// 2. If url is not valid, return error
// 3. If password is empty, generate password
// func newAccount(login, password, urlString string) (*account, error) {
// 	if login == "" {
// 		return nil, errors.New("EMPTY_LOGIN")
// 	}
	
// 	_, err := url.Parse(urlString)
// 	if err != nil {
// 		return nil, errors.New("INVALID_URL")
// 	}

// 	newAcc := &account{
// 		password: password,
// 		login:    login,
// 		url:      urlString,
// 	}
// 	if password == "" {
// 		newAcc.generatePassword(12)
// 	}
// 	return newAcc, nil
// }

func newAccountWithTimeStamp(login, password, urlString string) (*accountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}
	
	_, err := url.Parse(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &accountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		account: account{
			password: password,
			login:    login,
			url:      urlString,
		},
	}
	if password == "" {
		newAcc.generatePassword(12)
	}
	return newAcc, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

func main() {
	login := promptData("Input login")
	password := promptData("Input password")
	url := promptData("Input url")

	myAccount, err := newAccountWithTimeStamp(login, password, url)
	if err != nil {
		fmt.Println("Invalid url format")
		return
	}
	myAccount.outputPassword()
	fmt.Println(myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scan(&res)
	return res
}
