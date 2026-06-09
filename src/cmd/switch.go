package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// COMMAND LINE INTERFACE GOES HERE
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
	Run: func(cmd *cobra.Command, args []string) {
		performBranchSwitch(args[0])
	},
}

var switchAccountCmd = &cobra.Command{
	Use:   "account [username]",
	Short: "Switch to a specific user account",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		performAccountSwitch(args[0])
	},
}

// PERFORMERS GO HERE
func performAccountSwitch(username string) {
	fmt.Printf("Switched to account: %s\n", username)
}

func performBranchSwitch(branchName string) {
	fmt.Printf("Switched to branch: %s\n", branchName)
}