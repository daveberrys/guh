package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

const appDirName = "dev.pages.codedave.guh"
const accountsFileName = "accounts.json"

type Account struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	ClassicToken string `json:"classic_token"`
	Platform     string `json:"platform"`
}

func AccountsFilePath() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(configDir, appDirName, accountsFileName), nil
}

func LoadAccounts() ([]Account, error) {
	accountsFilePath, err := AccountsFilePath()
	if err != nil {
		return nil, err
	}

	contents, err := os.ReadFile(accountsFilePath)
	if errors.Is(err, os.ErrNotExist) {
		return []Account{}, nil
	}
	if err != nil {
		return nil, err
	}
	if len(contents) == 0 {
		return []Account{}, nil
	}

	var accounts []Account
	if err := json.Unmarshal(contents, &accounts); err != nil {
		return nil, fmt.Errorf("read accounts file: %w", err)
	}

	return accounts, nil
}

func SaveAccounts(accounts []Account) error {
	accountsFilePath, err := AccountsFilePath()
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(accountsFilePath), 0700); err != nil {
		return err
	}

	contents, err := json.MarshalIndent(accounts, "", "  ")
	if err != nil {
		return err
	}
	contents = append(contents, '\n')

	return os.WriteFile(accountsFilePath, contents, 0600)
}

func FindAccount(username string) (Account, error) {
	accounts, err := LoadAccounts()
	if err != nil {
		return Account{}, err
	}

	for _, account := range accounts {
		if account.Username == username {
			return account, nil
		}
	}

	return Account{}, fmt.Errorf("account %q not found", username)
}

func UpsertAccount(newAccount Account) error {
	accounts, err := LoadAccounts()
	if err != nil {
		return err
	}

	for index, account := range accounts {
		if account.Username == newAccount.Username {
			accounts[index] = newAccount
			return SaveAccounts(accounts)
		}
	}

	accounts = append(accounts, newAccount)
	return SaveAccounts(accounts)
}
