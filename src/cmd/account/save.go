package account

import (
	"fmt"
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

var saveCmd = &cobra.Command{
	Use:   "save [username] [email] [classicToken] [platform]",
	Short: "Save a new account",
	Args:  cobra.ExactArgs(4),
	RunE: func(c *cobra.Command, args []string) error {
		account := utils.Account{
			Username:     args[0],
			Email:        args[1],
			ClassicToken: args[2],
			Platform:     args[3],
		}
		if strings.TrimSpace(account.Username) == "" || strings.TrimSpace(account.Email) == "" || strings.TrimSpace(account.ClassicToken) == "" || strings.TrimSpace(account.Platform) == "" {
			return fmt.Errorf("username, email, classicToken, platform are required")
		}

		if err := utils.UpsertAccount(account); err != nil {
			return err
		}

		fmt.Printf("Saved account: %s\n", account.Username)
		return nil
	},
}
