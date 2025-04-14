package account

import (
	"errors"
	"math/rand/v2"
	"net/url"
	"time"
	"reflect"
	"fmt"
	"github.com/fatih/color"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890!@#$%^&*()")

type Account struct {
	login    string `json:"login" xml:"login"`
	password string `json:"password"`
	url      string `json:"url"`
}

type AccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	Account
}

func (acc *Account) OutputPassword() {
	color.Red("Login: %s\nPassword: %s\nUrl: %s\n", acc.login, acc.password, acc.url)
}

func (acc *Account) GeneratePassword(n int) {
	res := make([]rune, n)
	for i := range res {
		res[i] = letterRunes[rand.IntN(len(letterRunes))]
	}
	acc.password = string(res)
}

func NewAccountWithTimeStamp(login, password, urlString string) (*AccountWithTimeStamp, error) {
	if login == "" {
		return nil, errors.New("EMPTY_LOGIN")
	}

	_, err := url.Parse(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAcc := &AccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Account: Account{
			password: password,
			login:    login,
			url:      urlString,
		},
	}

    field, _ := reflect.TypeOf(newAcc).Elem().FieldByName("login")
    fmt.Println(string(field.Tag))

	if password == "" {
		newAcc.GeneratePassword(12)
	}
	return newAcc, nil
}
