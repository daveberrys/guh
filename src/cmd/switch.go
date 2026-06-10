package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/daveberrys/guh/src/utils"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(switchCmd)
	switchCmd.AddCommand(switchBranchCmd)
	switchCmd.AddCommand(switchAccountCmd)
}

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch branches or accounts",
}

var switchBranchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "Switch to a specific branch",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return performBranchSwitch(args[0])
	},
}

var switchAccountCmd = &cobra.Command{
	Use:   "account [username]",
	Short: "Switch to a specific user account",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return performAccountSwitch(args[0])
	},
}

func performAccountSwitch(username string) error {
	account, err := utils.FindAccount(username)
	if err != nil { return err }

	if err := utils.RunGitSequence(false,
		[]string{"config", "--global", "user.name", account.Username},
		[]string{"config", "--global", "user.email", account.Email},
	); err != nil { return err }

	homeDir, err := os.UserHomeDir()
	if err != nil { return err }

	credentials := fmt.Sprintf("https://%s:%s@github.com\n", account.Username, account.ClassicToken)
	if err := os.WriteFile(filepath.Join(homeDir, ".git-credentials"), []byte(credentials), 0600); err != nil { return err }

	fmt.Printf("Switched to account: %s\n", account.Username)
	return nil
}

func performBranchSwitch(branchName string) error {
	return utils.RunGitSequence(false, []string{"switch", branchName})
}
