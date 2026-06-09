package cmd

import (
	"fmt"
	"strings"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createBranchCmd)
	createCmd.AddCommand(createAccountCmd)
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create branches or accounts",
}

var createBranchCmd = &cobra.Command{
	Use:   "branch [branch]",
	Short: "Create and switch to a branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return utils.RunGitSequence([]string{"switch", "-c", args[0]})
	},
}

var createAccountCmd = &cobra.Command{
	Use:   "account [username] [email] [classicToken]",
	Short: "Create or update an account",
	Args:  cobra.ExactArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		account := utils.Account{
			Username:     args[0],
			Email:        args[1],
			ClassicToken: args[2],
		}
		if strings.TrimSpace(account.Username) == "" || strings.TrimSpace(account.Email) == "" || strings.TrimSpace(account.ClassicToken) == "" {
			return fmt.Errorf("username, email, and classicToken are required")
		}

		if err := utils.UpsertAccount(account); err != nil {
			return err
		}

		fmt.Printf("Saved account: %s\n", account.Username)
		return nil
	},
}
