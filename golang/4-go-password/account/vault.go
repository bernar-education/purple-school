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


type VaultWithDb struct {
    Vault
    db files.JsonDb
}


func (vault *Vault) ToByteSlice() ([]byte, error) {
    file, err := json.Marshal(vault)
    if err != nil {
        return nil, err
    }
    return file, nil
}


func NewVault(db *files.JsonDb) *VaultWithDb {
    file, err := db.Read()
    if err != nil {
        return &VaultWithDb{
            Vault: Vault{
                Accounts: []Account{},
                UpdatedAt: time.Now(),
            },
            db: *db,
        }
    }
    var vault Vault
    err = json.Unmarshal(file, &vault)
    if err != nil {
        color.Red("Can not unmarshal file data")
        return  &VaultWithDb{
            Vault: Vault{
                Accounts: []Account{},
                UpdatedAt: time.Now(),
            },
            db: *db,
        }
    }
    return &VaultWithDb{
        Vault: vault,
        db: *db,
    }
}


func (vault *VaultWithDb) FindAccountByUrl(url string) []Account {
    var accounts []Account
    for _, account := range vault.Vault.Accounts {
        isMatched := strings.Contains(account.Url, url)
        if isMatched {
            accounts = append(accounts, account)
        }
    }
    return accounts
}


func (vault *VaultWithDb) AddAccount(acc Account) {
    vault.Accounts = append(vault.Accounts, acc)
    vault.save()
}


func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
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


func (vault *VaultWithDb) save() {
    vault.UpdatedAt = time.Now()
    data, err := vault.Vault.ToByteSlice()
    if err != nil {
        color.Red("Can not convert to bytes")
    }
    vault.db.Write(data)
}
