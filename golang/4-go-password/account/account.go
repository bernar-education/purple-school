package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

type Account struct {
	Login       string      `json:"login"`
	Password    string      `json:"password"`
	Url         string      `json:"url"`
    CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

func (acc *Account) OutputPassword() {
	color.Red("Login: %s\nPassword: %s\nUrl: %s\n", acc.Login, acc.Password, acc.Url)
}


func (acc *Account) GeneratePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.Password = string(res)
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}

	_, err := url.Parse(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &Account{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
        Password: password,
        Login:    login,
        Url:      urlString,
	}

	if password == "" {
		newAcc.GeneratePassword(12)
	}
	return newAcc, nil
}
