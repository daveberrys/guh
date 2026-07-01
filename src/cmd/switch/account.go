package switchcmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daveberrys/guh/src/utils"
)

func performAccountSwitch(username string) error {
	account, err := utils.FindAccount(username)
	if err != nil { return err }

	if err := utils.RunGitSequence(false,
		[]string{"config", "--global", "user.name", account.Username},
		[]string{"config", "--global", "user.email", account.Email},
	); err != nil { return err }

	homeDir, err := os.UserHomeDir()
	if err != nil { return err }

	credentials := fmt.Sprintf("https://%s:%s@%s\n", account.Username, account.ClassicToken, account.Platform)
	if err := os.WriteFile(filepath.Join(homeDir, ".git-credentials"), []byte(credentials), 0600); err != nil { return err }

	fmt.Printf("Switched to account: %s\n", account.Username)
	return nil
}
