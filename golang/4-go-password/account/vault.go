package account

import (
	"encoding/json"
	"time"
	"strings"
	"github.com/fatih/color"
	"demo/password/files"
)

type Vault struct {
    Accounts   []Account   `json:"accounts"`
    UpdatedAt  time.Time   `json:"updatedAt"`
}


func (vault *Vault) ToByteSlice() ([]byte, error) {
    file, err := json.Marshal(vault)
    if err != nil {
        return nil, err
    }
    return file, nil
}


func NewVault() *Vault {
    db := files.NewJsonDb("data.json")
    file, err := db.Read()
    if err != nil {
        return &Vault{
            Accounts: []Account{},
            UpdatedAt: time.Now(),
        }
    }

    var vault Vault
    err = json.Unmarshal(file, &vault)
    if err != nil {
        color.Red("Can not unmarshal file data")
        return &Vault{
            Accounts: []Account{},
            UpdatedAt: time.Now(),
        }
    }
    return &vault
}


func (vault *Vault) FindAccountByUrl(url string) []Account {
    var accounts []Account
    for _, account := range vault.Accounts {
        isMatched := strings.Contains(account.Url, url)
        if isMatched {
            accounts = append(accounts, account)
        }
    }
    return accounts
}


func (vault *Vault) AddAccount(acc Account) {
    vault.Accounts = append(vault.Accounts, acc)
    vault.save()
}


func (vault * Vault) DeleteAccountByUrl(url string) bool {
    var accounts []Account
    isDeleted := false
    for _, account := range vault.Accounts {
        isMatched := strings.Contains(account.Url, url)
        if !isMatched {
            // vault.Accounts = append(vault.Accounts[:i], vault.Accounts[i+1:]...)
            accounts = append(accounts, account)
            continue
        }
        isDeleted = true
    }
    vault.Accounts = accounts
    vault.save()
    return isDeleted
}


func (vault *Vault) save() {
    vault.UpdatedAt = time.Now()
    data, err := vault.ToByteSlice()
    if err != nil {
        color.Red("Can not convert to bytes")
    }
    db := files.NewJsonDb("data.json")
    db.Write(data)
}
