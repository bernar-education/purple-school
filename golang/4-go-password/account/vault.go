package account

import (
	"demo/password/encrypter"
	"demo/password/output"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type ByteReader interface {
	Read() ([]byte, error)
}

type ByteWriter interface {
	Write([]byte)
}

type Db interface {
	ByteReader
	ByteWriter
}

type Vault struct {
	Accounts  []Account `json:"accounts"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db  Db
	enc encrypter.Encrypter
}

func (vault *Vault) ToByteSlice() ([]byte, error) {
	file, err := json.Marshal(vault)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewVault(db Db, enc encrypter.Encrypter) *VaultWithDb {
	file, err := db.Read()
	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	data := enc.Decrypt(file)
	var vault Vault
	err = json.Unmarshal(data, &vault)
	color.Cyan("Found %d accounts", len(vault.Accounts))
	if err != nil {
		output.PrintError("Can not unmarshal file data")
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []Account{},
				UpdatedAt: time.Now(),
			},
			db:  db,
			enc: enc,
		}
	}
	return &VaultWithDb{
		Vault: vault,
		db:    db,
		enc:   enc,
	}
}

func (vault *VaultWithDb) FindAccounts(str string, checker func(Account, string) bool) []Account {
	var accounts []Account
	for _, account := range vault.Vault.Accounts {
		isMatched := checker(account, str)
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
	encData := vault.enc.Encrypt(data)
	if err != nil {
		output.PrintError(err)
	}
	vault.db.Write(encData)
}
