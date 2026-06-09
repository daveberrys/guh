package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// COMMAND LINE INTERFACE GOES HERE
func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.AddCommand(accountSaveCmd)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save accounts,",
}

var accountSaveCmd = &cobra.Command{
	Use:   "account [username]",
	Short: "Save a specific user account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		performAccountSave(args[0])
	},
}

// PERFORMERS GO HERE
func performAccountSave(username string) {
	fmt.Printf("Saved account: %s\n", username)
}